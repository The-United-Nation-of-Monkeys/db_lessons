package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	connectErrorString        = "connect to database error: %v"
	migrateConnectErrorString = "connect for migrate database error: %v"
	migrateRunErrorString     = "run migrate error: %v"
)

type Config struct {
	Port uint16 `env:"POSTGRES_PORT" env-default:"5432"`
	Host string `env:"POSTGRES_HOST" env-default:"127.0.0.1"`

	Name     string `env:"POSTGRES_DB" env-default:"postgres"`
	User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`

	MaxConn int32 `env:"POSTGRES_MAX_CONN" env-default:"15"`
	MinConn int32 `env:"POSTGRES_MIN_CONN" env-default:"5"`
}

func Ping(ctx context.Context, cfg Config) bool {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return false
	}
	defer func() {
		if err = conn.Close(ctx); err != nil {
			log.Fatalf("close database conn error: %v", err)
		}
	}()

	err = conn.Ping(ctx)
	if err != nil {
		return false
	}
	return true
}

func New(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf(connectErrorString, err)
		return nil, err
	}

	poolConfig.MaxConns = cfg.MaxConn
	poolConfig.MinConns = cfg.MinConn

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf(connectErrorString, err)
		return nil, err
	}

	// Try to get migrations path from environment variable first
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		// Fallback to relative path calculation
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		migrationsPath = filepath.Join(basePath, "../../migrations")
	}
	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		))
	if err != nil {
		log.Fatalf(migrateConnectErrorString, err)
		return nil, err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf(migrateRunErrorString, err)
		return nil, err
	}

	return pool, nil
}

// QueryTracer логирует SQL запросы
type QueryTracer struct{}

func (qt *QueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	// Сохраняем время начала запроса в контексте
	startTime := time.Now()
	ctx = context.WithValue(ctx, "query_start_time", startTime)
	
	// Получаем пользователя из контекста
	dbUser := "unknown"
	if user, ok := ctx.Value("db_user").(string); ok {
		dbUser = user
	}
	
	if localLogger := logger.GetLoggerFromCtx(ctx); localLogger != nil {
		// Логируем начало запроса
		localLogger.Info(ctx, "SQL Query Start",
			zap.String("db_user", dbUser),
			zap.String("sql", truncateSQL(data.SQL, 200)),
			zap.Any("args", data.Args),
		)
	}
	return ctx
}

func (qt *QueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if localLogger := logger.GetLoggerFromCtx(ctx); localLogger != nil {
		dbUser := "unknown"
		if user, ok := ctx.Value("db_user").(string); ok {
			dbUser = user
		}
		
		// Получаем время начала из контекста
		var duration time.Duration
		if startTime, ok := ctx.Value("query_start_time").(time.Time); ok {
			duration = time.Since(startTime)
		}
		
		if data.Err != nil {
			localLogger.Error(ctx, "SQL Query Error",
				zap.String("db_user", dbUser),
				zap.Error(data.Err),
				zap.String("command_tag", data.CommandTag.String()),
				zap.Duration("duration", duration),
			)
		} else {
			localLogger.Info(ctx, "SQL Query End",
				zap.String("db_user", dbUser),
				zap.String("command_tag", data.CommandTag.String()),
				zap.Duration("duration", duration),
			)
		}
	}
}

func truncateSQL(sql string, maxLen int) string {
	if len(sql) <= maxLen {
		return sql
	}
	return sql[:maxLen] + "..."
}

func NewConn(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	config, err := pgx.ParseConfig(connString)
	if err != nil {
		log.Fatalf(connectErrorString, err)
		return nil, err
	}

	// Добавляем tracer для логирования SQL запросов
	if logger.GetLoggerFromCtx(ctx) != nil {
		config.Tracer = &QueryTracer{}
		// Сохраняем пользователя в контекст для логирования
		ctx = context.WithValue(ctx, "db_user", cfg.User)
	}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf(connectErrorString, err)
		return nil, err
	}
	return conn, nil
}

// RunMigrations применяет миграции к базе данных
// Использует указанного пользователя для подключения (обычно postgres для миграций)
func RunMigrations(ctx context.Context, cfg Config) error {
	// Try to get migrations path from environment variable first
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		// Fallback to relative path calculation
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		migrationsPath = filepath.Join(basePath, "../../migrations")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	m, err := migrate.New(
		"file://"+migrationsPath,
		connString,
	)
	if err != nil {
		log.Fatalf(migrateConnectErrorString, err)
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf(migrateRunErrorString, err)
		return err
	}

	log.Printf("Migrations applied successfully")
	return nil
}

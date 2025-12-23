package service

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type RoleConfig struct {
	User string
}

type BaseService struct {
	role        map[string]RoleConfig
	baseUser    string          
	originalCfg database.Config 
}

func NewBaseService(roleMap map[string]RoleConfig, cfg database.Config, baseUser string) *BaseService {
	if baseUser == "" {
		baseUser = "app_base_user" 
	}
	return &BaseService{
		role:        roleMap,
		baseUser:    baseUser,
		originalCfg: cfg,
	}
}

func (b *BaseService) GetRoleFromContext(ctx context.Context) string {
	if role, ok := ctx.Value("user_role").(string); ok && role != "" {
		return role
	}
	return ""
}

func (b *BaseService) GetDBConn(ctx context.Context, role string) (*pgx.Conn, error) {
	if role == "" {
		role = b.GetRoleFromContext(ctx)
	}

	cfg := b.originalCfg
	cfg.Password = "123" 

	var dbUser string
	if role == "" {
		dbUser = b.baseUser
		role = "base"
	} else if roleConfig, exists := b.role[role]; exists && roleConfig.User != "" {
		dbUser = roleConfig.User
	} else {
		dbUser = b.baseUser
		role = "base"
	}

	cfg.User = dbUser

	ctx = context.WithValue(ctx, "db_user", dbUser)
	ctx = context.WithValue(ctx, "db_role", role)

	if localLogger := logger.GetLoggerFromCtx(ctx); localLogger != nil {
		localLogger.Info(ctx, "Database connection",
			zap.String("db_user", dbUser),
			zap.String("role", role),
			zap.String("database", cfg.Name),
			zap.String("host", cfg.Host),
		)
	}

	dbConn, err := database.NewConn(ctx, cfg)
	if err != nil {
		if localLogger := logger.GetLoggerFromCtx(ctx); localLogger != nil {
			localLogger.Error(ctx, "Failed to connect to database",
				zap.String("db_user", dbUser),
				zap.String("role", role),
				zap.Error(err),
			)
		}
		return nil, err
	}

	return dbConn, nil
}

func (b *BaseService) BeginTx(ctx context.Context, role string) (*pgx.Conn, pgx.Tx, error) {
	conn, err := b.GetDBConn(ctx, role)
	if err != nil {
		return nil, nil, err
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		conn.Close(ctx)
		return nil, nil, err
	}

	return conn, tx, nil
}

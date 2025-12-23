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
	baseUser    string          // базовый пользователь по умолчанию
	originalCfg database.Config // оригинальная конфигурация
}

func NewBaseService(roleMap map[string]RoleConfig, cfg database.Config, baseUser string) *BaseService {
	if baseUser == "" {
		baseUser = "app_base_user" // значение по умолчанию
	}
	return &BaseService{
		role:        roleMap,
		baseUser:    baseUser,
		originalCfg: cfg,
	}
}

// GetRoleFromContext извлекает роль из контекста (из middleware)
func (b *BaseService) GetRoleFromContext(ctx context.Context) string {
	// Пытаемся получить роль из контекста Go (устанавливается в middleware через SetUserValue)
	if role, ok := ctx.Value("user_role").(string); ok && role != "" {
		return role
	}
	return ""
}

// GetDBConn возвращает соединение с БД для указанной роли
// Если role пустая, пытается получить роль из контекста
// Если роль не найдена, использует базового пользователя
func (b *BaseService) GetDBConn(ctx context.Context, role string) (*pgx.Conn, error) {
	// Если роль не указана, пытаемся получить из контекста
	if role == "" {
		role = b.GetRoleFromContext(ctx)
	}

	// Создаем копию конфигурации, чтобы не изменять оригинальную
	// ВАЖНО: НЕ используем cfg.User - только ролевых пользователей!
	cfg := b.originalCfg
	cfg.Password = "123" // Пароль для всех ролевых пользователей

	// Определяем пользователя для подключения
	var dbUser string
	if role == "" {
		// Если роли нет, используем базового пользователя
		dbUser = b.baseUser
		role = "base"
	} else if roleConfig, exists := b.role[role]; exists && roleConfig.User != "" {
		// Если роль есть в мапе и у неё указан пользователь
		dbUser = roleConfig.User
	} else {
		// Если роли нет в мапе или пользователь не указан, используем базового
		dbUser = b.baseUser
		role = "base"
	}

	cfg.User = dbUser

	// Сохраняем информацию о пользователе и роли в контекст для логирования
	ctx = context.WithValue(ctx, "db_user", dbUser)
	ctx = context.WithValue(ctx, "db_role", role)

	// Логируем подключение к БД
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

// BeginTx начинает транзакцию для указанной роли
// Если role пустая, пытается получить роль из контекста
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

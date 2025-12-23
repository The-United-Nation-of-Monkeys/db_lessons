package service

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"go.uber.org/zap"
)

type SQLExecuteServiceInterface interface {
	ExecuteSQL(ctx context.Context, data *dto.ExecuteSQLDTO) (*dto.ExecuteSQLResponseDTO, error)
}

type SQLExecuteService struct {
	baseService          *BaseService
	sqlExecuteRepository repository.SQLExecuteRepositoryInterface
}

func NewSQLExecuteService(baseService *BaseService, sqlExecuteRepository repository.SQLExecuteRepositoryInterface) *SQLExecuteService {
	return &SQLExecuteService{
		baseService:          baseService,
		sqlExecuteRepository: sqlExecuteRepository,
	}
}

func (s *SQLExecuteService) ExecuteSQL(ctx context.Context, data *dto.ExecuteSQLDTO) (*dto.ExecuteSQLResponseDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func ExecuteSQL")

	conn, err := s.baseService.GetDBConn(ctx, "admin")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	result, err := s.sqlExecuteRepository.ExecuteSQL(ctx, conn, data.Query)
	if err != nil {
		localLogger.Error(ctx, "execute sql error", zap.Error(err))
		return nil, exception.BadRequest("SQL execution error: " + err.Error())
	}

	localLogger.Info(ctx, "finish srv func ExecuteSQL")
	return result, nil
}


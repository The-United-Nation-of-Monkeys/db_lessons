package service

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"go.uber.org/zap"
)

type StatusTransactionServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStatusTransactionDTO) (*dto.StatusTransactionDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StatusTransactionDTO, error)
	GetAll(ctx context.Context) ([]*dto.StatusTransactionDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStatusTransactionDTO) (*dto.StatusTransactionDTO, error)
	Delete(ctx context.Context, id int) error
}

type StatusTransactionService struct {
	baseService *BaseService
	statusTransactionRepository repository.StatusTransactionRepositoryInterface
}

func NewStatusTransactionService(baseService *BaseService, statusTransactionRepository repository.StatusTransactionRepositoryInterface) *StatusTransactionService {
	return &StatusTransactionService{
		baseService: baseService,
		statusTransactionRepository: statusTransactionRepository,
	}
}

func (s *StatusTransactionService) Create(ctx context.Context, data *dto.CreateStatusTransactionDTO) (*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statustransaction, err := s.statusTransactionRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return statustransaction, nil
}

func (s *StatusTransactionService) GetByID(ctx context.Context, id int) (*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statustransaction, err := s.statusTransactionRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statustransaction not found")
			return nil, exception.NotFound("statustransaction not found")
		}
		localLogger.Error(ctx, "get statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return statustransaction, nil
}

func (s *StatusTransactionService) GetAll(ctx context.Context) ([]*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statustransactions, err := s.statusTransactionRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all statustransactions error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return statustransactions, nil
}

func (s *StatusTransactionService) Update(ctx context.Context, id int, data *dto.UpdateStatusTransactionDTO) (*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.statusTransactionRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statustransaction not found")
			return nil, exception.NotFound("statustransaction not found")
		}
		localLogger.Error(ctx, "get statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	statustransaction, err := s.statusTransactionRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return statustransaction, nil
}

func (s *StatusTransactionService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.statusTransactionRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statustransaction not found")
			return exception.NotFound("statustransaction not found")
		}
		localLogger.Error(ctx, "get statustransaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.statusTransactionRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete statustransaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

package service

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
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
	dbPool                      *pgxpool.Pool
	statusTransactionRepository repository.StatusTransactionRepositoryInterface
}

func NewStatusTransactionService(dbPool *pgxpool.Pool, statusTransactionRepository repository.StatusTransactionRepositoryInterface) *StatusTransactionService {
	return &StatusTransactionService{
		dbPool:                      dbPool,
		statusTransactionRepository: statusTransactionRepository,
	}
}

func (s *StatusTransactionService) Create(ctx context.Context, data *dto.CreateStatusTransactionDTO) (*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statustransaction, err := s.statusTransactionRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return statustransaction, nil
}

func (s *StatusTransactionService) GetByID(ctx context.Context, id int) (*dto.StatusTransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statustransaction, err := s.statusTransactionRepository.GetByID(ctx, tx, id)
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

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statustransactions, err := s.statusTransactionRepository.GetAll(ctx, tx)
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

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusTransactionRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statustransaction not found")
			return nil, exception.NotFound("statustransaction not found")
		}
		localLogger.Error(ctx, "get statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	statustransaction, err := s.statusTransactionRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update statustransaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return statustransaction, nil
}

func (s *StatusTransactionService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusTransactionRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statustransaction not found")
			return exception.NotFound("statustransaction not found")
		}
		localLogger.Error(ctx, "get statustransaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.statusTransactionRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete statustransaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

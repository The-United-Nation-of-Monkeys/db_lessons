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

type TransactionServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error)
	GetByID(ctx context.Context, id int) (*dto.TransactionDTO, error)
	GetAll(ctx context.Context) ([]*dto.TransactionDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error)
	Delete(ctx context.Context, id int) error
}

type TransactionService struct {
	dbPool                *pgxpool.Pool
	transactionRepository repository.TransactionRepositoryInterface
}

func NewTransactionService(dbPool *pgxpool.Pool, transactionRepository repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{
		dbPool:                dbPool,
		transactionRepository: transactionRepository,
	}
}

func (s *TransactionService) Create(ctx context.Context, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	transaction, err := s.transactionRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return transaction, nil
}

func (s *TransactionService) GetByID(ctx context.Context, id int) (*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	transaction, err := s.transactionRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "transaction not found")
			return nil, exception.NotFound("transaction not found")
		}
		localLogger.Error(ctx, "get transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return transaction, nil
}

func (s *TransactionService) GetAll(ctx context.Context) ([]*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	transactions, err := s.transactionRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all transactions error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return transactions, nil
}

func (s *TransactionService) Update(ctx context.Context, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.transactionRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "transaction not found")
			return nil, exception.NotFound("transaction not found")
		}
		localLogger.Error(ctx, "get transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	transaction, err := s.transactionRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return transaction, nil
}

func (s *TransactionService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.transactionRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "transaction not found")
			return exception.NotFound("transaction not found")
		}
		localLogger.Error(ctx, "get transaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.transactionRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete transaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

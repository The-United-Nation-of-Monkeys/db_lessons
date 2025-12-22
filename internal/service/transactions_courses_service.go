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

type TransactionsCoursesServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error)
	GetByID(ctx context.Context, transactionID, courseID int) (*dto.TransactionsCoursesDTO, error)
	GetAll(ctx context.Context) ([]*dto.TransactionsCoursesDTO, error)
	Delete(ctx context.Context, transactionID, courseID int) error
}

type TransactionsCoursesService struct {
	dbPool                       *pgxpool.Pool
	transactionsCoursesRepository repository.TransactionsCoursesRepositoryInterface
}

func NewTransactionsCoursesService(dbPool *pgxpool.Pool, transactionsCoursesRepository repository.TransactionsCoursesRepositoryInterface) *TransactionsCoursesService {
	return &TransactionsCoursesService{
		dbPool:                       dbPool,
		transactionsCoursesRepository: transactionsCoursesRepository,
	}
}

func (s *TransactionsCoursesService) Create(ctx context.Context, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.transactionsCoursesRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return relation, nil
}

func (s *TransactionsCoursesService) GetByID(ctx context.Context, transactionID, courseID int) (*dto.TransactionsCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.transactionsCoursesRepository.GetByID(ctx, tx, transactionID, courseID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return nil, exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return relation, nil
}

func (s *TransactionsCoursesService) GetAll(ctx context.Context) ([]*dto.TransactionsCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relations, err := s.transactionsCoursesRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *TransactionsCoursesService) Delete(ctx context.Context, transactionID, courseID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.transactionsCoursesRepository.GetByID(ctx, tx, transactionID, courseID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.transactionsCoursesRepository.Delete(ctx, tx, transactionID, courseID)
	if err != nil {
		localLogger.Error(ctx, "delete relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}


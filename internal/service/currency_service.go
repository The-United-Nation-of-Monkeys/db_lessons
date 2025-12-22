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

type CurrencyServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error)
	GetByID(ctx context.Context, id int) (*dto.CurrencyDTO, error)
	GetAll(ctx context.Context) ([]*dto.CurrencyDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateCurrencyDTO) (*dto.CurrencyDTO, error)
	Delete(ctx context.Context, id int) error
}

type CurrencyService struct {
	dbPool             *pgxpool.Pool
	currencyRepository repository.CurrencyRepositoryInterface
}

func NewCurrencyService(dbPool *pgxpool.Pool, currencyRepository repository.CurrencyRepositoryInterface) *CurrencyService {
	return &CurrencyService{
		dbPool:             dbPool,
		currencyRepository: currencyRepository,
	}
}

func (s *CurrencyService) Create(ctx context.Context, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	currency, err := s.currencyRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return currency, nil
}

func (s *CurrencyService) GetByID(ctx context.Context, id int) (*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	currency, err := s.currencyRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "currency not found")
			return nil, exception.NotFound("currency not found")
		}
		localLogger.Error(ctx, "get currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return currency, nil
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	currencys, err := s.currencyRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all currencys error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return currencys, nil
}

func (s *CurrencyService) Update(ctx context.Context, id int, data *dto.UpdateCurrencyDTO) (*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.currencyRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "currency not found")
			return nil, exception.NotFound("currency not found")
		}
		localLogger.Error(ctx, "get currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	currency, err := s.currencyRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return currency, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.currencyRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "currency not found")
			return exception.NotFound("currency not found")
		}
		localLogger.Error(ctx, "get currency error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.currencyRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete currency error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

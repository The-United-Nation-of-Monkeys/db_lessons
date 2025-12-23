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

type CurrencyServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error)
	GetByID(ctx context.Context, id int) (*dto.CurrencyDTO, error)
	GetAll(ctx context.Context) ([]*dto.CurrencyDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateCurrencyDTO) (*dto.CurrencyDTO, error)
	Delete(ctx context.Context, id int) error
}

type CurrencyService struct {
	baseService *BaseService
	currencyRepository repository.CurrencyRepositoryInterface
}

func NewCurrencyService(baseService *BaseService, currencyRepository repository.CurrencyRepositoryInterface) *CurrencyService {
	return &CurrencyService{
		baseService: baseService,
		currencyRepository: currencyRepository,
	}
}

func (s *CurrencyService) Create(ctx context.Context, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	currency, err := s.currencyRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return currency, nil
}

func (s *CurrencyService) GetByID(ctx context.Context, id int) (*dto.CurrencyDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	currency, err := s.currencyRepository.GetByID(ctx, conn, id)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	currencys, err := s.currencyRepository.GetAll(ctx, conn)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.currencyRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "currency not found")
			return nil, exception.NotFound("currency not found")
		}
		localLogger.Error(ctx, "get currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	currency, err := s.currencyRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update currency error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return currency, nil
}

func (s *CurrencyService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.currencyRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "currency not found")
			return exception.NotFound("currency not found")
		}
		localLogger.Error(ctx, "get currency error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.currencyRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete currency error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

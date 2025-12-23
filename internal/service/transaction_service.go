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

type TransactionServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error)
	GetByID(ctx context.Context, id int) (*dto.TransactionDTO, error)
	GetAll(ctx context.Context) ([]*dto.TransactionDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error)
	Delete(ctx context.Context, id int) error
	GetReportByParams(ctx context.Context, params *dto.TransactionReportRequestDTO) ([]*dto.TransactionReportDTO, error)
	BulkUpdateTransactionStatus(ctx context.Context, params *dto.BulkUpdateTransactionStatusDTO) error
}

type TransactionService struct {
	baseService *BaseService
	transactionRepository repository.TransactionRepositoryInterface
}

func NewTransactionService(baseService *BaseService, transactionRepository repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{
		baseService: baseService,
		transactionRepository: transactionRepository,
	}
}

func (s *TransactionService) Create(ctx context.Context, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	transaction, err := s.transactionRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return transaction, nil
}

func (s *TransactionService) GetByID(ctx context.Context, id int) (*dto.TransactionDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	transaction, err := s.transactionRepository.GetByID(ctx, conn, id)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	transactions, err := s.transactionRepository.GetAll(ctx, conn)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.transactionRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "transaction not found")
			return nil, exception.NotFound("transaction not found")
		}
		localLogger.Error(ctx, "get transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	transaction, err := s.transactionRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update transaction error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return transaction, nil
}

func (s *TransactionService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.transactionRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "transaction not found")
			return exception.NotFound("transaction not found")
		}
		localLogger.Error(ctx, "get transaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.transactionRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete transaction error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

func (s *TransactionService) GetReportByParams(ctx context.Context, params *dto.TransactionReportRequestDTO) ([]*dto.TransactionReportDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetReportByParams")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	reports, err := s.transactionRepository.GetReportByParams(ctx, conn, params)
	if err != nil {
		localLogger.Error(ctx, "get report by params error", zap.Error(err), zap.Any("params", params))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetReportByParams")
	return reports, nil
}

func (s *TransactionService) BulkUpdateTransactionStatus(ctx context.Context, params *dto.BulkUpdateTransactionStatusDTO) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func BulkUpdateTransactionStatus")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	err = s.transactionRepository.BulkUpdateTransactionStatus(ctx, conn, params)
	if err != nil {
		localLogger.Error(ctx, "bulk update transaction status error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func BulkUpdateTransactionStatus")
	return nil
}

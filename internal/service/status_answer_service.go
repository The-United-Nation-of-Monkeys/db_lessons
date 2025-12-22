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

type StatusAnswerServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStatusAnswerDTO) (*dto.StatusAnswerDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StatusAnswerDTO, error)
	GetAll(ctx context.Context) ([]*dto.StatusAnswerDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStatusAnswerDTO) (*dto.StatusAnswerDTO, error)
	Delete(ctx context.Context, id int) error
}

type StatusAnswerService struct {
	dbPool                 *pgxpool.Pool
	statusAnswerRepository repository.StatusAnswerRepositoryInterface
}

func NewStatusAnswerService(dbPool *pgxpool.Pool, statusAnswerRepository repository.StatusAnswerRepositoryInterface) *StatusAnswerService {
	return &StatusAnswerService{
		dbPool:                 dbPool,
		statusAnswerRepository: statusAnswerRepository,
	}
}

func (s *StatusAnswerService) Create(ctx context.Context, data *dto.CreateStatusAnswerDTO) (*dto.StatusAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statusanswer, err := s.statusAnswerRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create statusanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return statusanswer, nil
}

func (s *StatusAnswerService) GetByID(ctx context.Context, id int) (*dto.StatusAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statusanswer, err := s.statusAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statusanswer not found")
			return nil, exception.NotFound("statusanswer not found")
		}
		localLogger.Error(ctx, "get statusanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return statusanswer, nil
}

func (s *StatusAnswerService) GetAll(ctx context.Context) ([]*dto.StatusAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statusanswers, err := s.statusAnswerRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all statusanswers error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return statusanswers, nil
}

func (s *StatusAnswerService) Update(ctx context.Context, id int, data *dto.UpdateStatusAnswerDTO) (*dto.StatusAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statusanswer not found")
			return nil, exception.NotFound("statusanswer not found")
		}
		localLogger.Error(ctx, "get statusanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	statusanswer, err := s.statusAnswerRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update statusanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return statusanswer, nil
}

func (s *StatusAnswerService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statusanswer not found")
			return exception.NotFound("statusanswer not found")
		}
		localLogger.Error(ctx, "get statusanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.statusAnswerRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete statusanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

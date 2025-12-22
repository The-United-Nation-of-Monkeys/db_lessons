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

type StatusHomeworkServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StatusHomeworkDTO, error)
	GetAll(ctx context.Context) ([]*dto.StatusHomeworkDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	Delete(ctx context.Context, id int) error
}

type StatusHomeworkService struct {
	dbPool                   *pgxpool.Pool
	statusHomeworkRepository repository.StatusHomeworkRepositoryInterface
}

func NewStatusHomeworkService(dbPool *pgxpool.Pool, statusHomeworkRepository repository.StatusHomeworkRepositoryInterface) *StatusHomeworkService {
	return &StatusHomeworkService{
		dbPool:                   dbPool,
		statusHomeworkRepository: statusHomeworkRepository,
	}
}

func (s *StatusHomeworkService) Create(ctx context.Context, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statushomework, err := s.statusHomeworkRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return statushomework, nil
}

func (s *StatusHomeworkService) GetByID(ctx context.Context, id int) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statushomework, err := s.statusHomeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return nil, exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return statushomework, nil
}

func (s *StatusHomeworkService) GetAll(ctx context.Context) ([]*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	statushomeworks, err := s.statusHomeworkRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all statushomeworks error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return statushomeworks, nil
}

func (s *StatusHomeworkService) Update(ctx context.Context, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusHomeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return nil, exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	statushomework, err := s.statusHomeworkRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return statushomework, nil
}

func (s *StatusHomeworkService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.statusHomeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.statusHomeworkRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete statushomework error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

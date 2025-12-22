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

type HomeworkServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error)
	GetByID(ctx context.Context, id int) (*dto.HomeworkDTO, error)
	GetAll(ctx context.Context) ([]*dto.HomeworkDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error)
	Delete(ctx context.Context, id int) error
}

type HomeworkService struct {
	dbPool             *pgxpool.Pool
	homeworkRepository repository.HomeworkRepositoryInterface
}

func NewHomeworkService(dbPool *pgxpool.Pool, homeworkRepository repository.HomeworkRepositoryInterface) *HomeworkService {
	return &HomeworkService{
		dbPool:             dbPool,
		homeworkRepository: homeworkRepository,
	}
}

func (s *HomeworkService) Create(ctx context.Context, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homework, err := s.homeworkRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return homework, nil
}

func (s *HomeworkService) GetByID(ctx context.Context, id int) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homework, err := s.homeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return nil, exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return homework, nil
}

func (s *HomeworkService) GetAll(ctx context.Context) ([]*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homeworks, err := s.homeworkRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all homeworks error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return homeworks, nil
}

func (s *HomeworkService) Update(ctx context.Context, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.homeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return nil, exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	homework, err := s.homeworkRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return homework, nil
}

func (s *HomeworkService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.homeworkRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.homeworkRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete homework error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

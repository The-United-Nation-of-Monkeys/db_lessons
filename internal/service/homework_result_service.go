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

type HomeworkResultServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	GetByID(ctx context.Context, id int) (*dto.HomeworkResultDTO, error)
	GetAll(ctx context.Context) ([]*dto.HomeworkResultDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	Delete(ctx context.Context, id int) error
}

type HomeworkResultService struct {
	dbPool                   *pgxpool.Pool
	homeworkResultRepository repository.HomeworkResultRepositoryInterface
}

func NewHomeworkResultService(dbPool *pgxpool.Pool, homeworkResultRepository repository.HomeworkResultRepositoryInterface) *HomeworkResultService {
	return &HomeworkResultService{
		dbPool:                   dbPool,
		homeworkResultRepository: homeworkResultRepository,
	}
}

func (s *HomeworkResultService) Create(ctx context.Context, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homeworkresult, err := s.homeworkResultRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return homeworkresult, nil
}

func (s *HomeworkResultService) GetByID(ctx context.Context, id int) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homeworkresult, err := s.homeworkResultRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return nil, exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return homeworkresult, nil
}

func (s *HomeworkResultService) GetAll(ctx context.Context) ([]*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	homeworkresults, err := s.homeworkResultRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all homeworkresults error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return homeworkresults, nil
}

func (s *HomeworkResultService) Update(ctx context.Context, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.homeworkResultRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return nil, exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	homeworkresult, err := s.homeworkResultRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return homeworkresult, nil
}

func (s *HomeworkResultService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.homeworkResultRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.homeworkResultRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete homeworkresult error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

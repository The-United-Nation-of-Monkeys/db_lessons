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

type LessonHomeworksServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error)
	GetByID(ctx context.Context, lessonID, homeworkID int) (*dto.LessonHomeworksDTO, error)
	GetAll(ctx context.Context) ([]*dto.LessonHomeworksDTO, error)
	Delete(ctx context.Context, lessonID, homeworkID int) error
}

type LessonHomeworksService struct {
	dbPool                    *pgxpool.Pool
	lessonHomeworksRepository repository.LessonHomeworksRepositoryInterface
}

func NewLessonHomeworksService(dbPool *pgxpool.Pool, lessonHomeworksRepository repository.LessonHomeworksRepositoryInterface) *LessonHomeworksService {
	return &LessonHomeworksService{
		dbPool:                    dbPool,
		lessonHomeworksRepository: lessonHomeworksRepository,
	}
}

func (s *LessonHomeworksService) Create(ctx context.Context, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.lessonHomeworksRepository.Create(ctx, tx, data)
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

func (s *LessonHomeworksService) GetByID(ctx context.Context, lessonID, homeworkID int) (*dto.LessonHomeworksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.lessonHomeworksRepository.GetByID(ctx, tx, lessonID, homeworkID)
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

func (s *LessonHomeworksService) GetAll(ctx context.Context) ([]*dto.LessonHomeworksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relations, err := s.lessonHomeworksRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *LessonHomeworksService) Delete(ctx context.Context, lessonID, homeworkID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.lessonHomeworksRepository.GetByID(ctx, tx, lessonID, homeworkID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.lessonHomeworksRepository.Delete(ctx, tx, lessonID, homeworkID)
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


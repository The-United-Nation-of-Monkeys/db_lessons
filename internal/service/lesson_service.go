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

type LessonServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateLessonDTO) (*dto.LessonDTO, error)
	GetByID(ctx context.Context, id int) (*dto.LessonDTO, error)
	GetAll(ctx context.Context) ([]*dto.LessonDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateLessonDTO) (*dto.LessonDTO, error)
	Delete(ctx context.Context, id int) error
}

type LessonService struct {
	dbPool           *pgxpool.Pool
	lessonRepository repository.LessonRepositoryInterface
}

func NewLessonService(dbPool *pgxpool.Pool, lessonRepository repository.LessonRepositoryInterface) *LessonService {
	return &LessonService{
		dbPool:           dbPool,
		lessonRepository: lessonRepository,
	}
}

func (s *LessonService) Create(ctx context.Context, data *dto.CreateLessonDTO) (*dto.LessonDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	lesson, err := s.lessonRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create lesson error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return lesson, nil
}

func (s *LessonService) GetByID(ctx context.Context, id int) (*dto.LessonDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	lesson, err := s.lessonRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "lesson not found")
			return nil, exception.NotFound("lesson not found")
		}
		localLogger.Error(ctx, "get lesson error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return lesson, nil
}

func (s *LessonService) GetAll(ctx context.Context) ([]*dto.LessonDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	lessons, err := s.lessonRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all lessons error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return lessons, nil
}

func (s *LessonService) Update(ctx context.Context, id int, data *dto.UpdateLessonDTO) (*dto.LessonDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.lessonRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "lesson not found")
			return nil, exception.NotFound("lesson not found")
		}
		localLogger.Error(ctx, "get lesson error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	lesson, err := s.lessonRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update lesson error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return lesson, nil
}

func (s *LessonService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.lessonRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "lesson not found")
			return exception.NotFound("lesson not found")
		}
		localLogger.Error(ctx, "get lesson error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.lessonRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete lesson error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

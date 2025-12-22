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

type LessonsMaterialsServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error)
	GetByID(ctx context.Context, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error)
	GetAll(ctx context.Context) ([]*dto.LessonsMaterialsDTO, error)
	Delete(ctx context.Context, lessonID, materialID int) error
}

type LessonsMaterialsService struct {
	dbPool                      *pgxpool.Pool
	lessonsMaterialsRepository repository.LessonsMaterialsRepositoryInterface
}

func NewLessonsMaterialsService(dbPool *pgxpool.Pool, lessonsMaterialsRepository repository.LessonsMaterialsRepositoryInterface) *LessonsMaterialsService {
	return &LessonsMaterialsService{
		dbPool:                      dbPool,
		lessonsMaterialsRepository: lessonsMaterialsRepository,
	}
}

func (s *LessonsMaterialsService) Create(ctx context.Context, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.lessonsMaterialsRepository.Create(ctx, tx, data)
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

func (s *LessonsMaterialsService) GetByID(ctx context.Context, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relation, err := s.lessonsMaterialsRepository.GetByID(ctx, tx, lessonID, materialID)
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

func (s *LessonsMaterialsService) GetAll(ctx context.Context) ([]*dto.LessonsMaterialsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	relations, err := s.lessonsMaterialsRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *LessonsMaterialsService) Delete(ctx context.Context, lessonID, materialID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.lessonsMaterialsRepository.GetByID(ctx, tx, lessonID, materialID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.lessonsMaterialsRepository.Delete(ctx, tx, lessonID, materialID)
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


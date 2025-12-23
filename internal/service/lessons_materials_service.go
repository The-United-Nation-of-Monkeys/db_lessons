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

type LessonsMaterialsServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error)
	GetByID(ctx context.Context, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error)
	GetAll(ctx context.Context) ([]*dto.LessonsMaterialsDTO, error)
	Delete(ctx context.Context, lessonID, materialID int) error
}

type LessonsMaterialsService struct {
	baseService *BaseService
	lessonsMaterialsRepository repository.LessonsMaterialsRepositoryInterface
}

func NewLessonsMaterialsService(baseService *BaseService, lessonsMaterialsRepository repository.LessonsMaterialsRepositoryInterface) *LessonsMaterialsService {
	return &LessonsMaterialsService{
		baseService: baseService,
		lessonsMaterialsRepository: lessonsMaterialsRepository,
	}
}

func (s *LessonsMaterialsService) Create(ctx context.Context, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.lessonsMaterialsRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return relation, nil
}

func (s *LessonsMaterialsService) GetByID(ctx context.Context, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.lessonsMaterialsRepository.GetByID(ctx, conn, lessonID, materialID)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relations, err := s.lessonsMaterialsRepository.GetAll(ctx, conn)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.lessonsMaterialsRepository.GetByID(ctx, conn, lessonID, materialID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.lessonsMaterialsRepository.Delete(ctx, conn, lessonID, materialID)
	if err != nil {
		localLogger.Error(ctx, "delete relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}


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

type CourseLessonsServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error)
	GetByID(ctx context.Context, lessonID int) (*dto.CourseLessonsDTO, error)
	GetAll(ctx context.Context) ([]*dto.CourseLessonsDTO, error)
	Delete(ctx context.Context, lessonID int) error
}

type CourseLessonsService struct {
	baseService *BaseService
	courseLessonsRepository repository.CourseLessonsRepositoryInterface
}

func NewCourseLessonsService(baseService *BaseService, courseLessonsRepository repository.CourseLessonsRepositoryInterface) *CourseLessonsService {
	return &CourseLessonsService{
		baseService: baseService,
		courseLessonsRepository: courseLessonsRepository,
	}
}

func (s *CourseLessonsService) Create(ctx context.Context, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.courseLessonsRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return relation, nil
}

func (s *CourseLessonsService) GetByID(ctx context.Context, lessonID int) (*dto.CourseLessonsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.courseLessonsRepository.GetByID(ctx, conn, lessonID)
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

func (s *CourseLessonsService) GetAll(ctx context.Context) ([]*dto.CourseLessonsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relations, err := s.courseLessonsRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *CourseLessonsService) Delete(ctx context.Context, lessonID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.courseLessonsRepository.GetByID(ctx, conn, lessonID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.courseLessonsRepository.Delete(ctx, conn, lessonID)
	if err != nil {
		localLogger.Error(ctx, "delete relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}


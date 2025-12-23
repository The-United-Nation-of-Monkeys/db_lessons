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

type CourseServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCourseDTO) (*dto.CourseDTO, error)
	GetByID(ctx context.Context, id int) (*dto.CourseDTO, error)
	GetAll(ctx context.Context) ([]*dto.CourseDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateCourseDTO) (*dto.CourseDTO, error)
	Delete(ctx context.Context, id int) error
}

type CourseService struct {
	baseService *BaseService
	courseRepository repository.CourseRepositoryInterface
}

func NewCourseService(baseService *BaseService, courseRepository repository.CourseRepositoryInterface) *CourseService {
	return &CourseService{
		baseService: baseService,
		courseRepository: courseRepository,
	}
}

func (s *CourseService) Create(ctx context.Context, data *dto.CreateCourseDTO) (*dto.CourseDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	course, err := s.courseRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create course error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return course, nil
}

func (s *CourseService) GetByID(ctx context.Context, id int) (*dto.CourseDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	course, err := s.courseRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "course not found")
			return nil, exception.NotFound("course not found")
		}
		localLogger.Error(ctx, "get course error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return course, nil
}

func (s *CourseService) GetAll(ctx context.Context) ([]*dto.CourseDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	courses, err := s.courseRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all courses error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return courses, nil
}

func (s *CourseService) Update(ctx context.Context, id int, data *dto.UpdateCourseDTO) (*dto.CourseDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.courseRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "course not found")
			return nil, exception.NotFound("course not found")
		}
		localLogger.Error(ctx, "get course error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	course, err := s.courseRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update course error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return course, nil
}

func (s *CourseService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.courseRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "course not found")
			return exception.NotFound("course not found")
		}
		localLogger.Error(ctx, "get course error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.courseRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete course error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

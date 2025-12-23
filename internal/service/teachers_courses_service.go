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

type TeachersCoursesServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error)
	GetByID(ctx context.Context, courseID, teacherID int) (*dto.TeachersCoursesDTO, error)
	GetAll(ctx context.Context) ([]*dto.TeachersCoursesDTO, error)
	Delete(ctx context.Context, courseID, teacherID int) error
}

type TeachersCoursesService struct {
	baseService *BaseService
	teachersCoursesRepository repository.TeachersCoursesRepositoryInterface
}

func NewTeachersCoursesService(baseService *BaseService, teachersCoursesRepository repository.TeachersCoursesRepositoryInterface) *TeachersCoursesService {
	return &TeachersCoursesService{
		baseService: baseService,
		teachersCoursesRepository: teachersCoursesRepository,
	}
}

func (s *TeachersCoursesService) Create(ctx context.Context, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.teachersCoursesRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return relation, nil
}

func (s *TeachersCoursesService) GetByID(ctx context.Context, courseID, teacherID int) (*dto.TeachersCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.teachersCoursesRepository.GetByID(ctx, conn, courseID, teacherID)
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

func (s *TeachersCoursesService) GetAll(ctx context.Context) ([]*dto.TeachersCoursesDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relations, err := s.teachersCoursesRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *TeachersCoursesService) Delete(ctx context.Context, courseID, teacherID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.teachersCoursesRepository.GetByID(ctx, conn, courseID, teacherID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.teachersCoursesRepository.Delete(ctx, conn, courseID, teacherID)
	if err != nil {
		localLogger.Error(ctx, "delete relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}


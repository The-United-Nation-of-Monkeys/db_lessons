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

type TeacherServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error)
	GetByID(ctx context.Context, id int) (*dto.TeacherDTO, error)
	GetAll(ctx context.Context) ([]*dto.TeacherDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error)
	Delete(ctx context.Context, id int) error
}

type TeacherService struct {
	baseService       *BaseService
	teacherRepository repository.TeacherRepositoryInterface
}

func NewTeacherService(baseService *BaseService, teacherRepository repository.TeacherRepositoryInterface) *TeacherService {
	return &TeacherService{
		baseService:       baseService,
		teacherRepository: teacherRepository,
	}
}

func (s *TeacherService) Create(ctx context.Context, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	teacher, err := s.teacherRepository.Create(ctx, conn, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "create teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return teacher, nil
}

func (s *TeacherService) GetByID(ctx context.Context, id int) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	teacher, err := s.teacherRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return nil, exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return teacher, nil
}

func (s *TeacherService) GetAll(ctx context.Context) ([]*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	teachers, err := s.teacherRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all teachers error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return teachers, nil
}

func (s *TeacherService) Update(ctx context.Context, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.teacherRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return nil, exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	teacher, err := s.teacherRepository.Update(ctx, conn, id, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "update teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return teacher, nil
}

func (s *TeacherService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.teacherRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.teacherRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete teacher error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

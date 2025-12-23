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

type StudentServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStudentDTO) (*dto.StudentDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StudentDTO, error)
	GetAll(ctx context.Context) ([]*dto.StudentDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error)
	Delete(ctx context.Context, id int) error
}

type StudentService struct {
	baseService       *BaseService
	studentRepository repository.StudentRepositoryInterface
}

func NewStudentService(baseService *BaseService, studentRepository repository.StudentRepositoryInterface) *StudentService {
	return &StudentService{
		baseService:       baseService,
		studentRepository: studentRepository,
	}
}

func (s *StudentService) Create(ctx context.Context, data *dto.CreateStudentDTO) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	student, err := s.studentRepository.Create(ctx, conn, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "create student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return student, nil
}

func (s *StudentService) GetByID(ctx context.Context, id int) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	student, err := s.studentRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return nil, exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return student, nil
}

func (s *StudentService) GetAll(ctx context.Context) ([]*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	students, err := s.studentRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all students error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return students, nil
}

func (s *StudentService) Update(ctx context.Context, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.studentRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return nil, exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	student, err := s.studentRepository.Update(ctx, conn, id, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "update student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return student, nil
}

func (s *StudentService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.studentRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.studentRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete student error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

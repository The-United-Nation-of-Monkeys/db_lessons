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

type StudentAnswerServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StudentAnswerDTO, error)
	GetAll(ctx context.Context) ([]*dto.StudentAnswerDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	Delete(ctx context.Context, id int) error
}

type StudentAnswerService struct {
	baseService *BaseService
	studentAnswerRepository repository.StudentAnswerRepositoryInterface
}

func NewStudentAnswerService(baseService *BaseService, studentAnswerRepository repository.StudentAnswerRepositoryInterface) *StudentAnswerService {
	return &StudentAnswerService{
		baseService: baseService,
		studentAnswerRepository: studentAnswerRepository,
	}
}

func (s *StudentAnswerService) Create(ctx context.Context, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	studentanswer, err := s.studentAnswerRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return studentanswer, nil
}

func (s *StudentAnswerService) GetByID(ctx context.Context, id int) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	studentanswer, err := s.studentAnswerRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return nil, exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return studentanswer, nil
}

func (s *StudentAnswerService) GetAll(ctx context.Context) ([]*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	studentanswers, err := s.studentAnswerRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all studentanswers error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return studentanswers, nil
}

func (s *StudentAnswerService) Update(ctx context.Context, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.studentAnswerRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return nil, exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	studentanswer, err := s.studentAnswerRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return studentanswer, nil
}

func (s *StudentAnswerService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.studentAnswerRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.studentAnswerRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete studentanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

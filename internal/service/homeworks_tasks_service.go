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

type HomeworksTasksServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error)
	GetByID(ctx context.Context, taskID, homeworkID int) (*dto.HomeworksTasksDTO, error)
	GetAll(ctx context.Context) ([]*dto.HomeworksTasksDTO, error)
	Delete(ctx context.Context, taskID, homeworkID int) error
}

type HomeworksTasksService struct {
	baseService *BaseService
	homeworksTasksRepository repository.HomeworksTasksRepositoryInterface
}

func NewHomeworksTasksService(baseService *BaseService, homeworksTasksRepository repository.HomeworksTasksRepositoryInterface) *HomeworksTasksService {
	return &HomeworksTasksService{
		baseService: baseService,
		homeworksTasksRepository: homeworksTasksRepository,
	}
}

func (s *HomeworksTasksService) Create(ctx context.Context, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.homeworksTasksRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create relation error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return relation, nil
}

func (s *HomeworksTasksService) GetByID(ctx context.Context, taskID, homeworkID int) (*dto.HomeworksTasksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relation, err := s.homeworksTasksRepository.GetByID(ctx, conn, taskID, homeworkID)
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

func (s *HomeworksTasksService) GetAll(ctx context.Context) ([]*dto.HomeworksTasksDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	relations, err := s.homeworksTasksRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all relations error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return relations, nil
}

func (s *HomeworksTasksService) Delete(ctx context.Context, taskID, homeworkID int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.homeworksTasksRepository.GetByID(ctx, conn, taskID, homeworkID)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "relation not found")
			return exception.NotFound("relation not found")
		}
		localLogger.Error(ctx, "get relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.homeworksTasksRepository.Delete(ctx, conn, taskID, homeworkID)
	if err != nil {
		localLogger.Error(ctx, "delete relation error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}


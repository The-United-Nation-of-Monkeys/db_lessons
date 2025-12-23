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

type TaskServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTaskDTO) (*dto.TaskDTO, error)
	GetByID(ctx context.Context, id int) (*dto.TaskDTO, error)
	GetAll(ctx context.Context) ([]*dto.TaskDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateTaskDTO) (*dto.TaskDTO, error)
	Delete(ctx context.Context, id int) error
}

type TaskService struct {
	baseService *BaseService
	taskRepository repository.TaskRepositoryInterface
}

func NewTaskService(baseService *BaseService, taskRepository repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{
		baseService: baseService,
		taskRepository: taskRepository,
	}
}

func (s *TaskService) Create(ctx context.Context, data *dto.CreateTaskDTO) (*dto.TaskDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	task, err := s.taskRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create task error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return task, nil
}

func (s *TaskService) GetByID(ctx context.Context, id int) (*dto.TaskDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	task, err := s.taskRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "task not found")
			return nil, exception.NotFound("task not found")
		}
		localLogger.Error(ctx, "get task error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return task, nil
}

func (s *TaskService) GetAll(ctx context.Context) ([]*dto.TaskDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	tasks, err := s.taskRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all tasks error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return tasks, nil
}

func (s *TaskService) Update(ctx context.Context, id int, data *dto.UpdateTaskDTO) (*dto.TaskDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.taskRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "task not found")
			return nil, exception.NotFound("task not found")
		}
		localLogger.Error(ctx, "get task error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	task, err := s.taskRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update task error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return task, nil
}

func (s *TaskService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.taskRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "task not found")
			return exception.NotFound("task not found")
		}
		localLogger.Error(ctx, "get task error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.taskRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete task error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

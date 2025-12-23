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

type HomeworkServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error)
	GetByID(ctx context.Context, id int) (*dto.HomeworkDTO, error)
	GetAll(ctx context.Context) ([]*dto.HomeworkDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error)
	Delete(ctx context.Context, id int) error
}

type HomeworkService struct {
	baseService *BaseService
	homeworkRepository repository.HomeworkRepositoryInterface
}

func NewHomeworkService(baseService *BaseService, homeworkRepository repository.HomeworkRepositoryInterface) *HomeworkService {
	return &HomeworkService{
		baseService: baseService,
		homeworkRepository: homeworkRepository,
	}
}

func (s *HomeworkService) Create(ctx context.Context, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homework, err := s.homeworkRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return homework, nil
}

func (s *HomeworkService) GetByID(ctx context.Context, id int) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homework, err := s.homeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return nil, exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return homework, nil
}

func (s *HomeworkService) GetAll(ctx context.Context) ([]*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homeworks, err := s.homeworkRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all homeworks error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return homeworks, nil
}

func (s *HomeworkService) Update(ctx context.Context, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.homeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return nil, exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	homework, err := s.homeworkRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update homework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return homework, nil
}

func (s *HomeworkService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.homeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homework not found")
			return exception.NotFound("homework not found")
		}
		localLogger.Error(ctx, "get homework error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.homeworkRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete homework error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

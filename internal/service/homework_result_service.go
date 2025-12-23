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

type HomeworkResultServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	GetByID(ctx context.Context, id int) (*dto.HomeworkResultDTO, error)
	GetAll(ctx context.Context) ([]*dto.HomeworkResultDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	Delete(ctx context.Context, id int) error
}

type HomeworkResultService struct {
	baseService *BaseService
	homeworkResultRepository repository.HomeworkResultRepositoryInterface
}

func NewHomeworkResultService(baseService *BaseService, homeworkResultRepository repository.HomeworkResultRepositoryInterface) *HomeworkResultService {
	return &HomeworkResultService{
		baseService: baseService,
		homeworkResultRepository: homeworkResultRepository,
	}
}

func (s *HomeworkResultService) Create(ctx context.Context, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homeworkresult, err := s.homeworkResultRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return homeworkresult, nil
}

func (s *HomeworkResultService) GetByID(ctx context.Context, id int) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homeworkresult, err := s.homeworkResultRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return nil, exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return homeworkresult, nil
}

func (s *HomeworkResultService) GetAll(ctx context.Context) ([]*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	homeworkresults, err := s.homeworkResultRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all homeworkresults error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return homeworkresults, nil
}

func (s *HomeworkResultService) Update(ctx context.Context, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.homeworkResultRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return nil, exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	homeworkresult, err := s.homeworkResultRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update homeworkresult error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return homeworkresult, nil
}

func (s *HomeworkResultService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.homeworkResultRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "homeworkresult not found")
			return exception.NotFound("homeworkresult not found")
		}
		localLogger.Error(ctx, "get homeworkresult error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.homeworkResultRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete homeworkresult error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

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

type StatusHomeworkServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StatusHomeworkDTO, error)
	GetAll(ctx context.Context) ([]*dto.StatusHomeworkDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	Delete(ctx context.Context, id int) error
}

type StatusHomeworkService struct {
	baseService *BaseService
	statusHomeworkRepository repository.StatusHomeworkRepositoryInterface
}

func NewStatusHomeworkService(baseService *BaseService, statusHomeworkRepository repository.StatusHomeworkRepositoryInterface) *StatusHomeworkService {
	return &StatusHomeworkService{
		baseService: baseService,
		statusHomeworkRepository: statusHomeworkRepository,
	}
}

func (s *StatusHomeworkService) Create(ctx context.Context, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statushomework, err := s.statusHomeworkRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return statushomework, nil
}

func (s *StatusHomeworkService) GetByID(ctx context.Context, id int) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statushomework, err := s.statusHomeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return nil, exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return statushomework, nil
}

func (s *StatusHomeworkService) GetAll(ctx context.Context) ([]*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	statushomeworks, err := s.statusHomeworkRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all statushomeworks error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return statushomeworks, nil
}

func (s *StatusHomeworkService) Update(ctx context.Context, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.statusHomeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return nil, exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	statushomework, err := s.statusHomeworkRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update statushomework error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return statushomework, nil
}

func (s *StatusHomeworkService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.statusHomeworkRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "statushomework not found")
			return exception.NotFound("statushomework not found")
		}
		localLogger.Error(ctx, "get statushomework error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.statusHomeworkRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete statushomework error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

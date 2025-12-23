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

type LevelServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateLevelDTO) (*dto.LevelDTO, error)
	GetByID(ctx context.Context, id int) (*dto.LevelDTO, error)
	GetAll(ctx context.Context) ([]*dto.LevelDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error)
	Delete(ctx context.Context, id int) error
}

type LevelService struct {
	baseService *BaseService
	levelRepository repository.LevelRepositoryInterface
}

func NewLevelService(baseService *BaseService, levelRepository repository.LevelRepositoryInterface) *LevelService {
	return &LevelService{
		baseService: baseService,
		levelRepository: levelRepository,
	}
}

func (s *LevelService) Create(ctx context.Context, data *dto.CreateLevelDTO) (*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	level, err := s.levelRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return level, nil
}

func (s *LevelService) GetByID(ctx context.Context, id int) (*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	level, err := s.levelRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "level not found")
			return nil, exception.NotFound("level not found")
		}
		localLogger.Error(ctx, "get level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return level, nil
}

func (s *LevelService) GetAll(ctx context.Context) ([]*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	levels, err := s.levelRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all levels error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return levels, nil
}

func (s *LevelService) Update(ctx context.Context, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.levelRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "level not found")
			return nil, exception.NotFound("level not found")
		}
		localLogger.Error(ctx, "get level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	level, err := s.levelRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return level, nil
}

func (s *LevelService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.levelRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "level not found")
			return exception.NotFound("level not found")
		}
		localLogger.Error(ctx, "get level error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.levelRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete level error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

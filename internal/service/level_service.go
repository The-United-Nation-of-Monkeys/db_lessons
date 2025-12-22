package service

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
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
	dbPool          *pgxpool.Pool
	levelRepository repository.LevelRepositoryInterface
}

func NewLevelService(dbPool *pgxpool.Pool, levelRepository repository.LevelRepositoryInterface) *LevelService {
	return &LevelService{
		dbPool:          dbPool,
		levelRepository: levelRepository,
	}
}

func (s *LevelService) Create(ctx context.Context, data *dto.CreateLevelDTO) (*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	level, err := s.levelRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return level, nil
}

func (s *LevelService) GetByID(ctx context.Context, id int) (*dto.LevelDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	level, err := s.levelRepository.GetByID(ctx, tx, id)
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

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	levels, err := s.levelRepository.GetAll(ctx, tx)
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

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.levelRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "level not found")
			return nil, exception.NotFound("level not found")
		}
		localLogger.Error(ctx, "get level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	level, err := s.levelRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update level error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return level, nil
}

func (s *LevelService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.levelRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "level not found")
			return exception.NotFound("level not found")
		}
		localLogger.Error(ctx, "get level error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.levelRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete level error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

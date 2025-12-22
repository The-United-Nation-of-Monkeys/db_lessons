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

type MaterialServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateMaterialDTO) (*dto.MaterialDTO, error)
	GetByID(ctx context.Context, id int) (*dto.MaterialDTO, error)
	GetAll(ctx context.Context) ([]*dto.MaterialDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateMaterialDTO) (*dto.MaterialDTO, error)
	Delete(ctx context.Context, id int) error
}

type MaterialService struct {
	dbPool             *pgxpool.Pool
	materialRepository repository.MaterialRepositoryInterface
}

func NewMaterialService(dbPool *pgxpool.Pool, materialRepository repository.MaterialRepositoryInterface) *MaterialService {
	return &MaterialService{
		dbPool:             dbPool,
		materialRepository: materialRepository,
	}
}

func (s *MaterialService) Create(ctx context.Context, data *dto.CreateMaterialDTO) (*dto.MaterialDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	material, err := s.materialRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create material error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return material, nil
}

func (s *MaterialService) GetByID(ctx context.Context, id int) (*dto.MaterialDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	material, err := s.materialRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "material not found")
			return nil, exception.NotFound("material not found")
		}
		localLogger.Error(ctx, "get material error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return material, nil
}

func (s *MaterialService) GetAll(ctx context.Context) ([]*dto.MaterialDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	materials, err := s.materialRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all materials error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return materials, nil
}

func (s *MaterialService) Update(ctx context.Context, id int, data *dto.UpdateMaterialDTO) (*dto.MaterialDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.materialRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "material not found")
			return nil, exception.NotFound("material not found")
		}
		localLogger.Error(ctx, "get material error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	material, err := s.materialRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update material error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return material, nil
}

func (s *MaterialService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.materialRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "material not found")
			return exception.NotFound("material not found")
		}
		localLogger.Error(ctx, "get material error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.materialRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete material error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

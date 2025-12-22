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

type SubcategoryServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	GetByID(ctx context.Context, id int) (*dto.SubcategoryDTO, error)
	GetAll(ctx context.Context) ([]*dto.SubcategoryDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	Delete(ctx context.Context, id int) error
}

type SubcategoryService struct {
	dbPool                *pgxpool.Pool
	subcategoryRepository repository.SubcategoryRepositoryInterface
}

func NewSubcategoryService(dbPool *pgxpool.Pool, subcategoryRepository repository.SubcategoryRepositoryInterface) *SubcategoryService {
	return &SubcategoryService{
		dbPool:                dbPool,
		subcategoryRepository: subcategoryRepository,
	}
}

func (s *SubcategoryService) Create(ctx context.Context, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	subcategory, err := s.subcategoryRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return subcategory, nil
}

func (s *SubcategoryService) GetByID(ctx context.Context, id int) (*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	subcategory, err := s.subcategoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "subcategory not found")
			return nil, exception.NotFound("subcategory not found")
		}
		localLogger.Error(ctx, "get subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return subcategory, nil
}

func (s *SubcategoryService) GetAll(ctx context.Context) ([]*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	subcategorys, err := s.subcategoryRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all subcategorys error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return subcategorys, nil
}

func (s *SubcategoryService) Update(ctx context.Context, id int, data *dto.UpdateSubcategoryDTO) (*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.subcategoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "subcategory not found")
			return nil, exception.NotFound("subcategory not found")
		}
		localLogger.Error(ctx, "get subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	subcategory, err := s.subcategoryRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return subcategory, nil
}

func (s *SubcategoryService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.subcategoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "subcategory not found")
			return exception.NotFound("subcategory not found")
		}
		localLogger.Error(ctx, "get subcategory error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.subcategoryRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete subcategory error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

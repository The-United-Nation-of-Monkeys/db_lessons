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

type CategoryServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error)
	GetByID(ctx context.Context, id int) (*dto.CategoryDTO, error)
	GetAll(ctx context.Context) ([]*dto.CategoryDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error)
	Delete(ctx context.Context, id int) error
}

type CategoryService struct {
	dbPool             *pgxpool.Pool
	categoryRepository repository.CategoryRepositoryInterface
}

func NewCategoryService(dbPool *pgxpool.Pool, categoryRepository repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		dbPool:             dbPool,
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	category, err := s.categoryRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return category, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id int) (*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	category, err := s.categoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "category not found")
			return nil, exception.NotFound("category not found")
		}
		localLogger.Error(ctx, "get category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return category, nil
}

func (s *CategoryService) GetAll(ctx context.Context) ([]*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	categories, err := s.categoryRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all categories error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return categories, nil
}

func (s *CategoryService) Update(ctx context.Context, id int, data *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.categoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "category not found")
			return nil, exception.NotFound("category not found")
		}
		localLogger.Error(ctx, "get category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	category, err := s.categoryRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return category, nil
}

func (s *CategoryService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.categoryRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "category not found")
			return exception.NotFound("category not found")
		}
		localLogger.Error(ctx, "get category error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.categoryRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete category error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

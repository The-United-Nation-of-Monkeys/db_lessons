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

type CategoryServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error)
	GetByID(ctx context.Context, id int) (*dto.CategoryDTO, error)
	GetAll(ctx context.Context) ([]*dto.CategoryDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error)
	Delete(ctx context.Context, id int) error
}

type CategoryService struct {
	baseService *BaseService
	categoryRepository repository.CategoryRepositoryInterface
}

func NewCategoryService(baseService *BaseService, categoryRepository repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		baseService: baseService,
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(ctx context.Context, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	// Получаем соединение с БД (роль будет получена из контекста)
	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	category, err := s.categoryRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return category, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id int) (*dto.CategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	// Получаем соединение с БД (роль будет получена из контекста)
	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	category, err := s.categoryRepository.GetByID(ctx, conn, id)
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

	// Получаем соединение с БД (роль будет получена из контекста)
	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	categories, err := s.categoryRepository.GetAll(ctx, conn)
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

	// Получаем соединение с БД (роль будет получена из контекста)
	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.categoryRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "category not found")
			return nil, exception.NotFound("category not found")
		}
		localLogger.Error(ctx, "get category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	category, err := s.categoryRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update category error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return category, nil
}

func (s *CategoryService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	// Получаем соединение с БД (роль будет получена из контекста)
	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "get db conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.categoryRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "category not found")
			return exception.NotFound("category not found")
		}
		localLogger.Error(ctx, "get category error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.categoryRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete category error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

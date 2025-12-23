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

type SubcategoryServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	GetByID(ctx context.Context, id int) (*dto.SubcategoryDTO, error)
	GetAll(ctx context.Context) ([]*dto.SubcategoryDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	Delete(ctx context.Context, id int) error
}

type SubcategoryService struct {
	baseService *BaseService
	subcategoryRepository repository.SubcategoryRepositoryInterface
}

func NewSubcategoryService(baseService *BaseService, subcategoryRepository repository.SubcategoryRepositoryInterface) *SubcategoryService {
	return &SubcategoryService{
		baseService: baseService,
		subcategoryRepository: subcategoryRepository,
	}
}

func (s *SubcategoryService) Create(ctx context.Context, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	subcategory, err := s.subcategoryRepository.Create(ctx, conn, data)
	if err != nil {
		localLogger.Error(ctx, "create subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return subcategory, nil
}

func (s *SubcategoryService) GetByID(ctx context.Context, id int) (*dto.SubcategoryDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	subcategory, err := s.subcategoryRepository.GetByID(ctx, conn, id)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	subcategorys, err := s.subcategoryRepository.GetAll(ctx, conn)
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

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.subcategoryRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "subcategory not found")
			return nil, exception.NotFound("subcategory not found")
		}
		localLogger.Error(ctx, "get subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	subcategory, err := s.subcategoryRepository.Update(ctx, conn, id, data)
	if err != nil {
		localLogger.Error(ctx, "update subcategory error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return subcategory, nil
}

func (s *SubcategoryService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer conn.Close(ctx)

	_, err = s.subcategoryRepository.GetByID(ctx, conn, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "subcategory not found")
			return exception.NotFound("subcategory not found")
		}
		localLogger.Error(ctx, "get subcategory error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.subcategoryRepository.Delete(ctx, conn, id)
	if err != nil {
		localLogger.Error(ctx, "delete subcategory error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

package service

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"go.uber.org/zap"
)

type StudentCategoryStatsServiceInterface interface {
	GetAll(ctx context.Context) ([]*dto.StudentCategoryStatsDTO, error)
	GetByStudentID(ctx context.Context, studentID int) ([]*dto.StudentCategoryStatsDTO, error)
	GetByCategoryID(ctx context.Context, categoryID int) ([]*dto.StudentCategoryStatsDTO, error)
}

type StudentCategoryStatsService struct {
	baseService *BaseService
	studentCategoryStatsRepository repository.StudentCategoryStatsRepositoryInterface
}

func NewStudentCategoryStatsService(baseService *BaseService, studentCategoryStatsRepository repository.StudentCategoryStatsRepositoryInterface) *StudentCategoryStatsService {
	return &StudentCategoryStatsService{
		baseService: baseService,
		studentCategoryStatsRepository: studentCategoryStatsRepository,
	}
}

func (s *StudentCategoryStatsService) GetAll(ctx context.Context) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	stats, err := s.studentCategoryStatsRepository.GetAll(ctx, conn)
	if err != nil {
		localLogger.Error(ctx, "get all stats error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return stats, nil
}

func (s *StudentCategoryStatsService) GetByStudentID(ctx context.Context, studentID int) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByStudentID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	stats, err := s.studentCategoryStatsRepository.GetByStudentID(ctx, conn, studentID)
	if err != nil {
		localLogger.Error(ctx, "get stats by student id error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByStudentID")
	return stats, nil
}

func (s *StudentCategoryStatsService) GetByCategoryID(ctx context.Context, categoryID int) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByCategoryID")

	conn, err := s.baseService.GetDBConn(ctx, "")
	if err != nil {
		localLogger.Error(ctx, "begin conn error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer conn.Close(ctx)

	stats, err := s.studentCategoryStatsRepository.GetByCategoryID(ctx, conn, categoryID)
	if err != nil {
		localLogger.Error(ctx, "get stats by category id error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByCategoryID")
	return stats, nil
}


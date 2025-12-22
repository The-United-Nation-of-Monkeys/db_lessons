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

type StudentCategoryStatsServiceInterface interface {
	GetAll(ctx context.Context) ([]*dto.StudentCategoryStatsDTO, error)
	GetByStudentID(ctx context.Context, studentID int) ([]*dto.StudentCategoryStatsDTO, error)
	GetByCategoryID(ctx context.Context, categoryID int) ([]*dto.StudentCategoryStatsDTO, error)
}

type StudentCategoryStatsService struct {
	dbPool                          *pgxpool.Pool
	studentCategoryStatsRepository repository.StudentCategoryStatsRepositoryInterface
}

func NewStudentCategoryStatsService(dbPool *pgxpool.Pool, studentCategoryStatsRepository repository.StudentCategoryStatsRepositoryInterface) *StudentCategoryStatsService {
	return &StudentCategoryStatsService{
		dbPool:                          dbPool,
		studentCategoryStatsRepository: studentCategoryStatsRepository,
	}
}

func (s *StudentCategoryStatsService) GetAll(ctx context.Context) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	stats, err := s.studentCategoryStatsRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all stats error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return stats, nil
}

func (s *StudentCategoryStatsService) GetByStudentID(ctx context.Context, studentID int) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByStudentID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	stats, err := s.studentCategoryStatsRepository.GetByStudentID(ctx, tx, studentID)
	if err != nil {
		localLogger.Error(ctx, "get stats by student id error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByStudentID")
	return stats, nil
}

func (s *StudentCategoryStatsService) GetByCategoryID(ctx context.Context, categoryID int) ([]*dto.StudentCategoryStatsDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByCategoryID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	stats, err := s.studentCategoryStatsRepository.GetByCategoryID(ctx, tx, categoryID)
	if err != nil {
		localLogger.Error(ctx, "get stats by category id error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByCategoryID")
	return stats, nil
}


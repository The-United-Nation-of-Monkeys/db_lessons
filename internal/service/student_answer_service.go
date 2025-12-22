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

type StudentAnswerServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StudentAnswerDTO, error)
	GetAll(ctx context.Context) ([]*dto.StudentAnswerDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	Delete(ctx context.Context, id int) error
}

type StudentAnswerService struct {
	dbPool                  *pgxpool.Pool
	studentAnswerRepository repository.StudentAnswerRepositoryInterface
}

func NewStudentAnswerService(dbPool *pgxpool.Pool, studentAnswerRepository repository.StudentAnswerRepositoryInterface) *StudentAnswerService {
	return &StudentAnswerService{
		dbPool:                  dbPool,
		studentAnswerRepository: studentAnswerRepository,
	}
}

func (s *StudentAnswerService) Create(ctx context.Context, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	studentanswer, err := s.studentAnswerRepository.Create(ctx, tx, data)
	if err != nil {
		localLogger.Error(ctx, "create studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return studentanswer, nil
}

func (s *StudentAnswerService) GetByID(ctx context.Context, id int) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	studentanswer, err := s.studentAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return nil, exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return studentanswer, nil
}

func (s *StudentAnswerService) GetAll(ctx context.Context) ([]*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	studentanswers, err := s.studentAnswerRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all studentanswers error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return studentanswers, nil
}

func (s *StudentAnswerService) Update(ctx context.Context, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.studentAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return nil, exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	studentanswer, err := s.studentAnswerRepository.Update(ctx, tx, id, data)
	if err != nil {
		localLogger.Error(ctx, "update studentanswer error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return studentanswer, nil
}

func (s *StudentAnswerService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.studentAnswerRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "studentanswer not found")
			return exception.NotFound("studentanswer not found")
		}
		localLogger.Error(ctx, "get studentanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.studentAnswerRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete studentanswer error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

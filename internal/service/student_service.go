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

type StudentServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateStudentDTO) (*dto.StudentDTO, error)
	GetByID(ctx context.Context, id int) (*dto.StudentDTO, error)
	GetAll(ctx context.Context) ([]*dto.StudentDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error)
	Delete(ctx context.Context, id int) error
}

type StudentService struct {
	dbPool            *pgxpool.Pool
	studentRepository repository.StudentRepositoryInterface
}

func NewStudentService(dbPool *pgxpool.Pool, studentRepository repository.StudentRepositoryInterface) *StudentService {
	return &StudentService{
		dbPool:            dbPool,
		studentRepository: studentRepository,
	}
}

func (s *StudentService) Create(ctx context.Context, data *dto.CreateStudentDTO) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	student, err := s.studentRepository.Create(ctx, tx, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "create student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return student, nil
}

func (s *StudentService) GetByID(ctx context.Context, id int) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	student, err := s.studentRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return nil, exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return student, nil
}

func (s *StudentService) GetAll(ctx context.Context) ([]*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	students, err := s.studentRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all students error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return students, nil
}

func (s *StudentService) Update(ctx context.Context, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.studentRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return nil, exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	student, err := s.studentRepository.Update(ctx, tx, id, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "update student error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return student, nil
}

func (s *StudentService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.studentRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "student not found")
			return exception.NotFound("student not found")
		}
		localLogger.Error(ctx, "get student error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.studentRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete student error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

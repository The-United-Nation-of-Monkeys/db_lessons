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

type TeacherServiceInterface interface {
	Create(ctx context.Context, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error)
	GetByID(ctx context.Context, id int) (*dto.TeacherDTO, error)
	GetAll(ctx context.Context) ([]*dto.TeacherDTO, error)
	Update(ctx context.Context, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error)
	Delete(ctx context.Context, id int) error
}

type TeacherService struct {
	dbPool            *pgxpool.Pool
	teacherRepository repository.TeacherRepositoryInterface
}

func NewTeacherService(dbPool *pgxpool.Pool, teacherRepository repository.TeacherRepositoryInterface) *TeacherService {
	return &TeacherService{
		dbPool:            dbPool,
		teacherRepository: teacherRepository,
	}
}

func (s *TeacherService) Create(ctx context.Context, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Create")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	teacher, err := s.teacherRepository.Create(ctx, tx, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "create teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Create")
	return teacher, nil
}

func (s *TeacherService) GetByID(ctx context.Context, id int) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetByID")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	teacher, err := s.teacherRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return nil, exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetByID")
	return teacher, nil
}

func (s *TeacherService) GetAll(ctx context.Context) ([]*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func GetAll")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	teachers, err := s.teacherRepository.GetAll(ctx, tx)
	if err != nil {
		localLogger.Error(ctx, "get all teachers error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func GetAll")
	return teachers, nil
}

func (s *TeacherService) Update(ctx context.Context, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Update")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return nil, exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.teacherRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return nil, exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	teacher, err := s.teacherRepository.Update(ctx, tx, id, data)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeDuplicate {
			localLogger.Info(ctx, "duplicate email")
			return nil, exception.BadRequest("email already exists")
		}
		localLogger.Error(ctx, "update teacher error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return nil, exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Update")
	return teacher, nil
}

func (s *TeacherService) Delete(ctx context.Context, id int) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	localLogger.Info(ctx, "start srv func Delete")

	tx, err := s.dbPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "begin tx error", zap.Error(err))
		return exception.InternalServerError()
	}
	defer database.RollbackTx(ctx, tx)

	_, err = s.teacherRepository.GetByID(ctx, tx, id)
	if err != nil {
		pgErr := database.ValidatePgxError(err)
		if pgErr != nil && pgErr.Type == database.TypeNoRows {
			localLogger.Info(ctx, "teacher not found")
			return exception.NotFound("teacher not found")
		}
		localLogger.Error(ctx, "get teacher error", zap.Error(err))
		return exception.InternalServerError()
	}

	err = s.teacherRepository.Delete(ctx, tx, id)
	if err != nil {
		localLogger.Error(ctx, "delete teacher error", zap.Error(err))
		return exception.InternalServerError()
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit error", zap.Error(err))
		return exception.InternalServerError()
	}

	localLogger.Info(ctx, "finish srv func Delete")
	return nil
}

package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// HomeworksTasksRepository
type HomeworksTasksRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) (*dto.HomeworksTasksDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworksTasksDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) error
}

type HomeworksTasksRepository struct{}

func NewHomeworksTasksRepository() *HomeworksTasksRepository {
	return &HomeworksTasksRepository{}
}

func (r *HomeworksTasksRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error) {
	query := `INSERT INTO homeworks_tasks (task_id, homework_id) VALUES ($1, $2) RETURNING task_id, homework_id`
	var relation dto.HomeworksTasksDTO
	err := tx.QueryRow(ctx, query, data.TaskID, data.HomeworkID).Scan(&relation.TaskID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *HomeworksTasksRepository) GetByID(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) (*dto.HomeworksTasksDTO, error) {
	query := `SELECT task_id, homework_id FROM homeworks_tasks WHERE task_id = $1 AND homework_id = $2`
	var relation dto.HomeworksTasksDTO
	err := tx.QueryRow(ctx, query, taskID, homeworkID).Scan(&relation.TaskID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *HomeworksTasksRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworksTasksDTO, error) {
	query := `SELECT task_id, homework_id FROM homeworks_tasks`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.HomeworksTasksDTO
	for rows.Next() {
		var relation dto.HomeworksTasksDTO
		if err := rows.Scan(&relation.TaskID, &relation.HomeworkID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *HomeworksTasksRepository) Delete(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) error {
	query := `DELETE FROM homeworks_tasks WHERE task_id = $1 AND homework_id = $2`
	_, err := tx.Exec(ctx, query, taskID, homeworkID)
	return err
}


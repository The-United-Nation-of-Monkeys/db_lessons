package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type TaskRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTaskDTO) (*dto.TaskDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.TaskDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TaskDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateTaskDTO) (*dto.TaskDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTaskDTO) (*dto.TaskDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO task (type, description, right_answer, points, level_id, category_id, subcategory_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) 
			  RETURNING task_id, type, description, right_answer, points, level_id, category_id, subcategory_id`
	var task dto.TaskDTO
	err = tx.QueryRow(ctx, query, data.Type, data.Description, data.RightAnswer, data.Points, data.LevelID, data.CategoryID, data.SubcategoryID).
		Scan(&task.TaskID, &task.Type, &task.Description, &task.RightAnswer, &task.Points, &task.LevelID, &task.CategoryID, &task.SubcategoryID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.TaskDTO, error) {
	query := `SELECT task_id, type, description, right_answer, points, level_id, category_id, subcategory_id FROM task WHERE task_id = $1`
	var task dto.TaskDTO
	err := conn.QueryRow(ctx, query, id).Scan(&task.TaskID, &task.Type, &task.Description, &task.RightAnswer, &task.Points, &task.LevelID, &task.CategoryID, &task.SubcategoryID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TaskDTO, error) {
	query := `SELECT task_id, type, description, right_answer, points, level_id, category_id, subcategory_id FROM task`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*dto.TaskDTO
	for rows.Next() {
		var task dto.TaskDTO
		if err := rows.Scan(&task.TaskID, &task.Type, &task.Description, &task.RightAnswer, &task.Points, &task.LevelID, &task.CategoryID, &task.SubcategoryID); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateTaskDTO) (*dto.TaskDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE task SET 
			  type = COALESCE($1, type),
			  description = COALESCE($2, description),
			  right_answer = COALESCE($3, right_answer),
			  points = COALESCE($4, points),
			  level_id = COALESCE($5, level_id),
			  category_id = COALESCE($6, category_id),
			  subcategory_id = COALESCE($7, subcategory_id)
			  WHERE task_id = $8
			  RETURNING task_id, type, description, right_answer, points, level_id, category_id, subcategory_id`
	var task dto.TaskDTO
	err = tx.QueryRow(ctx, query, data.Type, data.Description, data.RightAnswer, data.Points, data.LevelID, data.CategoryID, data.SubcategoryID, id).
		Scan(&task.TaskID, &task.Type, &task.Description, &task.RightAnswer, &task.Points, &task.LevelID, &task.CategoryID, &task.SubcategoryID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM task WHERE task_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

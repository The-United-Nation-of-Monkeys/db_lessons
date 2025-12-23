package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type HomeworkResultRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.HomeworkResultDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.HomeworkResultDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type HomeworkResultRepository struct{}

func NewHomeworkResultRepository() *HomeworkResultRepository {
	return &HomeworkResultRepository{}
}

func (r *HomeworkResultRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	query := `INSERT INTO homework_result (student_answer_id, student_id, task_id, status_homework_id, points) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING homework_result_id, student_answer_id, student_id, task_id, status_homework_id, points`
	var result dto.HomeworkResultDTO
	err := conn.QueryRow(ctx, query, data.StudentAnswerID, data.StudentID, data.TaskID, data.StatusHomeworkID, data.Points).
		Scan(&result.HomeworkResultID, &result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.HomeworkResultDTO, error) {
	query := `SELECT homework_result_id, student_answer_id, student_id, task_id, status_homework_id, points FROM homework_result WHERE homework_result_id = $1`
	var result dto.HomeworkResultDTO
	err := conn.QueryRow(ctx, query, id).Scan(&result.HomeworkResultID, &result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.HomeworkResultDTO, error) {
	query := `SELECT homework_result_id, student_answer_id, student_id, task_id, status_homework_id, points FROM homework_result`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*dto.HomeworkResultDTO
	for rows.Next() {
		var result dto.HomeworkResultDTO
		if err := rows.Scan(&result.HomeworkResultID, &result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.StatusHomeworkID, &result.Points); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (r *HomeworkResultRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	query := `UPDATE homework_result SET 
			  student_answer_id = COALESCE($1, student_answer_id),
			  student_id = COALESCE($2, student_id),
			  task_id = COALESCE($3, task_id),
			  status_homework_id = COALESCE($4, status_homework_id),
			  points = COALESCE($5, points)
			  WHERE homework_result_id = $6
			  RETURNING homework_result_id, student_answer_id, student_id, task_id, status_homework_id, points`
	var result dto.HomeworkResultDTO
	err := conn.QueryRow(ctx, query, data.StudentAnswerID, data.StudentID, data.TaskID, data.StatusHomeworkID, data.Points, id).
		Scan(&result.HomeworkResultID, &result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	query := `DELETE FROM homework_result WHERE homework_result_id = $1`
	_, err := conn.Exec(ctx, query, id)
	return err
}

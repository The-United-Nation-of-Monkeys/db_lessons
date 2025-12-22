package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type HomeworkResultRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.HomeworkResultDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworkResultDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type HomeworkResultRepository struct{}

func NewHomeworkResultRepository() *HomeworkResultRepository {
	return &HomeworkResultRepository{}
}

func (r *HomeworkResultRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	query := `INSERT INTO homework_result (student_id, task_id, answer, status_homework_id, points) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING student_answer_id, student_id, task_id, answer, status_homework_id, points`
	var result dto.HomeworkResultDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.TaskID, data.Answer, data.StatusHomeworkID, data.Points).
		Scan(&result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.Answer, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.HomeworkResultDTO, error) {
	query := `SELECT student_answer_id, student_id, task_id, answer, status_homework_id, points FROM homework_result WHERE student_answer_id = $1`
	var result dto.HomeworkResultDTO
	err := tx.QueryRow(ctx, query, id).Scan(&result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.Answer, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworkResultDTO, error) {
	query := `SELECT student_answer_id, student_id, task_id, answer, status_homework_id, points FROM homework_result`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*dto.HomeworkResultDTO
	for rows.Next() {
		var result dto.HomeworkResultDTO
		if err := rows.Scan(&result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.Answer, &result.StatusHomeworkID, &result.Points); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (r *HomeworkResultRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateHomeworkResultDTO) (*dto.HomeworkResultDTO, error) {
	query := `UPDATE homework_result SET 
			  student_id = COALESCE($1, student_id),
			  task_id = COALESCE($2, task_id),
			  answer = COALESCE($3, answer),
			  status_homework_id = COALESCE($4, status_homework_id),
			  points = COALESCE($5, points)
			  WHERE student_answer_id = $6
			  RETURNING student_answer_id, student_id, task_id, answer, status_homework_id, points`
	var result dto.HomeworkResultDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.TaskID, data.Answer, data.StatusHomeworkID, data.Points, id).
		Scan(&result.StudentAnswerID, &result.StudentID, &result.TaskID, &result.Answer, &result.StatusHomeworkID, &result.Points)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *HomeworkResultRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM homework_result WHERE student_answer_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

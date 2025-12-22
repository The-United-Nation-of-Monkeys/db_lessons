package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type StudentAnswerRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.StudentAnswerDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.StudentAnswerDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type StudentAnswerRepository struct{}

func NewStudentAnswerRepository() *StudentAnswerRepository {
	return &StudentAnswerRepository{}
}

func (r *StudentAnswerRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	query := `INSERT INTO student_answer (student_id, task_id, answer, status_answer_id) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING student_answer_id, student_id, task_id, answer, status_answer_id`
	var answer dto.StudentAnswerDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.TaskID, data.Answer, data.StatusAnswerID).
		Scan(&answer.StudentAnswerID, &answer.StudentID, &answer.TaskID, &answer.Answer, &answer.StatusAnswerID)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *StudentAnswerRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.StudentAnswerDTO, error) {
	query := `SELECT student_answer_id, student_id, task_id, answer, status_answer_id FROM student_answer WHERE student_answer_id = $1`
	var answer dto.StudentAnswerDTO
	err := tx.QueryRow(ctx, query, id).Scan(&answer.StudentAnswerID, &answer.StudentID, &answer.TaskID, &answer.Answer, &answer.StatusAnswerID)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *StudentAnswerRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.StudentAnswerDTO, error) {
	query := `SELECT student_answer_id, student_id, task_id, answer, status_answer_id FROM student_answer`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []*dto.StudentAnswerDTO
	for rows.Next() {
		var answer dto.StudentAnswerDTO
		if err := rows.Scan(&answer.StudentAnswerID, &answer.StudentID, &answer.TaskID, &answer.Answer, &answer.StatusAnswerID); err != nil {
			return nil, err
		}
		answers = append(answers, &answer)
	}
	return answers, nil
}

func (r *StudentAnswerRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateStudentAnswerDTO) (*dto.StudentAnswerDTO, error) {
	query := `UPDATE student_answer SET 
			  student_id = COALESCE($1, student_id),
			  task_id = COALESCE($2, task_id),
			  answer = COALESCE($3, answer),
			  status_answer_id = COALESCE($4, status_answer_id)
			  WHERE student_answer_id = $5
			  RETURNING student_answer_id, student_id, task_id, answer, status_answer_id`
	var answer dto.StudentAnswerDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.TaskID, data.Answer, data.StatusAnswerID, id).
		Scan(&answer.StudentAnswerID, &answer.StudentID, &answer.TaskID, &answer.Answer, &answer.StatusAnswerID)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *StudentAnswerRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM student_answer WHERE student_answer_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

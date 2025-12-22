package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// LessonHomeworksRepository
type LessonHomeworksRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) (*dto.LessonHomeworksDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonHomeworksDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) error
}

type LessonHomeworksRepository struct{}

func NewLessonHomeworksRepository() *LessonHomeworksRepository {
	return &LessonHomeworksRepository{}
}

func (r *LessonHomeworksRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error) {
	query := `INSERT INTO lesson_homeworks (lesson_id, homework_id) VALUES ($1, $2) RETURNING lesson_id, homework_id`
	var relation dto.LessonHomeworksDTO
	err := tx.QueryRow(ctx, query, data.LessonID, data.HomeworkID).Scan(&relation.LessonID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonHomeworksRepository) GetByID(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) (*dto.LessonHomeworksDTO, error) {
	query := `SELECT lesson_id, homework_id FROM lesson_homeworks WHERE lesson_id = $1 AND homework_id = $2`
	var relation dto.LessonHomeworksDTO
	err := tx.QueryRow(ctx, query, lessonID, homeworkID).Scan(&relation.LessonID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonHomeworksRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonHomeworksDTO, error) {
	query := `SELECT lesson_id, homework_id FROM lesson_homeworks`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.LessonHomeworksDTO
	for rows.Next() {
		var relation dto.LessonHomeworksDTO
		if err := rows.Scan(&relation.LessonID, &relation.HomeworkID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *LessonHomeworksRepository) Delete(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) error {
	query := `DELETE FROM lesson_homeworks WHERE lesson_id = $1 AND homework_id = $2`
	_, err := tx.Exec(ctx, query, lessonID, homeworkID)
	return err
}


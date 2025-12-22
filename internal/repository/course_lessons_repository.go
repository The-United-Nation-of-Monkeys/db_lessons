package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// CourseLessonsRepository
type CourseLessonsRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, lessonID int) (*dto.CourseLessonsDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.CourseLessonsDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, lessonID int) error
}

type CourseLessonsRepository struct{}

func NewCourseLessonsRepository() *CourseLessonsRepository {
	return &CourseLessonsRepository{}
}

func (r *CourseLessonsRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error) {
	query := `INSERT INTO course_lessons (lesson_id, course_id) VALUES ($1, $2) RETURNING lesson_id, course_id`
	var relation dto.CourseLessonsDTO
	err := tx.QueryRow(ctx, query, data.LessonID, data.CourseID).Scan(&relation.LessonID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *CourseLessonsRepository) GetByID(ctx context.Context, tx pgx.Tx, lessonID int) (*dto.CourseLessonsDTO, error) {
	query := `SELECT lesson_id, course_id FROM course_lessons WHERE lesson_id = $1`
	var relation dto.CourseLessonsDTO
	err := tx.QueryRow(ctx, query, lessonID).Scan(&relation.LessonID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *CourseLessonsRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.CourseLessonsDTO, error) {
	query := `SELECT lesson_id, course_id FROM course_lessons`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.CourseLessonsDTO
	for rows.Next() {
		var relation dto.CourseLessonsDTO
		if err := rows.Scan(&relation.LessonID, &relation.CourseID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *CourseLessonsRepository) Delete(ctx context.Context, tx pgx.Tx, lessonID int) error {
	query := `DELETE FROM course_lessons WHERE lesson_id = $1`
	_, err := tx.Exec(ctx, query, lessonID)
	return err
}


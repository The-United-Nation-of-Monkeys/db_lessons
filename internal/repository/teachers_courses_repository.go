package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// TeachersCoursesRepository
type TeachersCoursesRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, courseID, teacherID int) (*dto.TeachersCoursesDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeachersCoursesDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, courseID, teacherID int) error
}

type TeachersCoursesRepository struct{}

func NewTeachersCoursesRepository() *TeachersCoursesRepository {
	return &TeachersCoursesRepository{}
}

func (r *TeachersCoursesRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error) {
	query := `INSERT INTO teachers_courses (course_id, teacher_id) VALUES ($1, $2) RETURNING course_id, teacher_id`
	var relation dto.TeachersCoursesDTO
	err := tx.QueryRow(ctx, query, data.CourseID, data.TeacherID).Scan(&relation.CourseID, &relation.TeacherID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TeachersCoursesRepository) GetByID(ctx context.Context, tx pgx.Tx, courseID, teacherID int) (*dto.TeachersCoursesDTO, error) {
	query := `SELECT course_id, teacher_id FROM teachers_courses WHERE course_id = $1 AND teacher_id = $2`
	var relation dto.TeachersCoursesDTO
	err := tx.QueryRow(ctx, query, courseID, teacherID).Scan(&relation.CourseID, &relation.TeacherID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TeachersCoursesRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeachersCoursesDTO, error) {
	query := `SELECT course_id, teacher_id FROM teachers_courses`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.TeachersCoursesDTO
	for rows.Next() {
		var relation dto.TeachersCoursesDTO
		if err := rows.Scan(&relation.CourseID, &relation.TeacherID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *TeachersCoursesRepository) Delete(ctx context.Context, tx pgx.Tx, courseID, teacherID int) error {
	query := `DELETE FROM teachers_courses WHERE course_id = $1 AND teacher_id = $2`
	_, err := tx.Exec(ctx, query, courseID, teacherID)
	return err
}


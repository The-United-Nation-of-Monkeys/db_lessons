package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type CourseRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCourseDTO) (*dto.CourseDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CourseDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CourseDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCourseDTO) (*dto.CourseDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type CourseRepository struct{}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

func (r *CourseRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCourseDTO) (*dto.CourseDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO course (name, description, category_id, start_date, end_date, price, currency_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) 
			  RETURNING course_id, name, description, category_id, start_date, end_date, price, currency_id`
	var course dto.CourseDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.CategoryID, data.StartDate, data.EndDate, data.Price, data.CurrencyID).
		Scan(&course.CourseID, &course.Name, &course.Description, &course.CategoryID, &course.StartDate, &course.EndDate, &course.Price, &course.CurrencyID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CourseDTO, error) {
	query := `SELECT course_id, name, description, category_id, start_date, end_date, price, currency_id FROM course WHERE course_id = $1`
	var course dto.CourseDTO
	err := conn.QueryRow(ctx, query, id).Scan(&course.CourseID, &course.Name, &course.Description, &course.CategoryID, &course.StartDate, &course.EndDate, &course.Price, &course.CurrencyID)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CourseDTO, error) {
	query := `SELECT course_id, name, description, category_id, start_date, end_date, price, currency_id FROM course`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*dto.CourseDTO
	for rows.Next() {
		var course dto.CourseDTO
		if err := rows.Scan(&course.CourseID, &course.Name, &course.Description, &course.CategoryID, &course.StartDate, &course.EndDate, &course.Price, &course.CurrencyID); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (r *CourseRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCourseDTO) (*dto.CourseDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE course SET 
			  name = COALESCE($1, name),
			  description = COALESCE($2, description),
			  category_id = COALESCE($3, category_id),
			  start_date = COALESCE($4, start_date),
			  end_date = COALESCE($5, end_date),
			  price = COALESCE($6, price),
			  currency_id = COALESCE($7, currency_id)
			  WHERE course_id = $8
			  RETURNING course_id, name, description, category_id, start_date, end_date, price, currency_id`
	var course dto.CourseDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.CategoryID, data.StartDate, data.EndDate, data.Price, data.CurrencyID, id).
		Scan(&course.CourseID, &course.Name, &course.Description, &course.CategoryID, &course.StartDate, &course.EndDate, &course.Price, &course.CurrencyID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM course WHERE course_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

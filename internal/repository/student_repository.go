package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type StudentRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateStudentDTO) (*dto.StudentDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.StudentDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.StudentDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateStudentDTO) (*dto.StudentDTO, error) {
	query := `INSERT INTO student (name, surname, email, password) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING student_id, name, surname, email, password`
	var student dto.StudentDTO
	err := tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, data.Password).
		Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.StudentDTO, error) {
	query := `SELECT student_id, name, surname, email, password FROM student WHERE student_id = $1`
	var student dto.StudentDTO
	err := tx.QueryRow(ctx, query, id).Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.StudentDTO, error) {
	query := `SELECT student_id, name, surname, email, password FROM student`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*dto.StudentDTO
	for rows.Next() {
		var student dto.StudentDTO
		if err := rows.Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password); err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (r *StudentRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error) {
	query := `UPDATE student SET 
			  name = COALESCE($1, name),
			  surname = COALESCE($2, surname),
			  email = COALESCE($3, email),
			  password = COALESCE($4, password)
			  WHERE student_id = $5
			  RETURNING student_id, name, surname, email, password`
	var student dto.StudentDTO
	err := tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, data.Password, id).
		Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM student WHERE student_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

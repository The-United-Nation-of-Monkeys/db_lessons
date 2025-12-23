package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type StudentRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStudentDTO) (*dto.StudentDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StudentDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StudentDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStudentDTO) (*dto.StudentDTO, error) {
	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO student (name, surname, email, password) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING student_id, name, surname, email, password`
	var student dto.StudentDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, string(hashedPassword)).
		Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StudentDTO, error) {
	query := `SELECT student_id, name, surname, email, password FROM student WHERE student_id = $1`
	var student dto.StudentDTO
	err := conn.QueryRow(ctx, query, id).Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StudentDTO, error) {
	query := `SELECT student_id, name, surname, email, password FROM student`
	rows, err := conn.Query(ctx, query)
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

func (r *StudentRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStudentDTO) (*dto.StudentDTO, error) {
	var hashedPassword *string
	if data.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashedStr := string(hashed)
		hashedPassword = &hashedStr
	}

	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE student SET 
			  name = COALESCE($1, name),
			  surname = COALESCE($2, surname),
			  email = COALESCE($3, email),
			  password = COALESCE($4, password)
			  WHERE student_id = $5
			  RETURNING student_id, name, surname, email, password`
	var student dto.StudentDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, hashedPassword, id).
		Scan(&student.StudentID, &student.Name, &student.Surname, &student.Email, &student.Password)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM student WHERE student_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

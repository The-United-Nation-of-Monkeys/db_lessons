package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type TeacherRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.TeacherDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeacherDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type TeacherRepository struct{}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{}
}

func (r *TeacherRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeacherDTO) (*dto.TeacherDTO, error) {
	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO teacher (name, surname, email, password) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING teacher_id, name, surname, email, password`
	var teacher dto.TeacherDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, string(hashedPassword)).
		Scan(&teacher.TeacherID, &teacher.Name, &teacher.Surname, &teacher.Email, &teacher.Password)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.TeacherDTO, error) {
	query := `SELECT teacher_id, name, surname, email, password FROM teacher WHERE teacher_id = $1`
	var teacher dto.TeacherDTO
	err := tx.QueryRow(ctx, query, id).Scan(&teacher.TeacherID, &teacher.Name, &teacher.Surname, &teacher.Email, &teacher.Password)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeacherDTO, error) {
	query := `SELECT teacher_id, name, surname, email, password FROM teacher`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*dto.TeacherDTO
	for rows.Next() {
		var teacher dto.TeacherDTO
		if err := rows.Scan(&teacher.TeacherID, &teacher.Name, &teacher.Surname, &teacher.Email, &teacher.Password); err != nil {
			return nil, err
		}
		teachers = append(teachers, &teacher)
	}
	return teachers, nil
}

func (r *TeacherRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateTeacherDTO) (*dto.TeacherDTO, error) {
	var hashedPassword *string
	if data.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		hashedStr := string(hashed)
		hashedPassword = &hashedStr
	}

	query := `UPDATE teacher SET 
			  name = COALESCE($1, name),
			  surname = COALESCE($2, surname),
			  email = COALESCE($3, email),
			  password = COALESCE($4, password)
			  WHERE teacher_id = $5
			  RETURNING teacher_id, name, surname, email, password`
	var teacher dto.TeacherDTO
	err := tx.QueryRow(ctx, query, data.Name, data.Surname, data.Email, hashedPassword, id).
		Scan(&teacher.TeacherID, &teacher.Name, &teacher.Surname, &teacher.Email, &teacher.Password)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM teacher WHERE teacher_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

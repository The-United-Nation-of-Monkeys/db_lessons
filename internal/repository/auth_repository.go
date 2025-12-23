package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryInterface interface {
	LoginStudent(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error)
	LoginTeacher(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error)
	LoginAdmin(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error)
	GetUserByID(ctx context.Context, tx pgx.Tx, userID int, role string) (*dto.AuthResponseDTO, error)
}

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *AuthRepository) LoginStudent(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error) {
	query := `SELECT student_id, name, surname, email, password FROM student WHERE email = $1`
	
	var studentID int
	var name, surname, emailDB, passwordHash string
	
	err := tx.QueryRow(ctx, query, email).Scan(&studentID, &name, &surname, &emailDB, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !checkPasswordHash(password, passwordHash) {
		return nil, errors.New("invalid email or password")
	}

	return &dto.AuthResponseDTO{
		UserID: studentID,
		Role:   "student",
		Name:   name,
		Surname: surname,
		Email:  emailDB,
	}, nil
}

func (r *AuthRepository) LoginTeacher(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error) {
	query := `SELECT teacher_id, name, surname, email, password FROM teacher WHERE email = $1`
	
	var teacherID int
	var name, surname, emailDB, passwordHash string
	
	err := tx.QueryRow(ctx, query, email).Scan(&teacherID, &name, &surname, &emailDB, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !checkPasswordHash(password, passwordHash) {
		return nil, errors.New("invalid email or password")
	}

	return &dto.AuthResponseDTO{
		UserID: teacherID,
		Role:   "teacher",
		Name:   name,
		Surname: surname,
		Email:  emailDB,
	}, nil
}

func (r *AuthRepository) LoginAdmin(ctx context.Context, tx pgx.Tx, email, password string) (*dto.AuthResponseDTO, error) {
	query := `SELECT admin_id, name, surname, email, password FROM admin WHERE email = $1`
	
	var adminID int
	var name, surname, emailDB, passwordHash string
	
	err := tx.QueryRow(ctx, query, email).Scan(&adminID, &name, &surname, &emailDB, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !checkPasswordHash(password, passwordHash) {
		return nil, errors.New("invalid email or password")
	}

	return &dto.AuthResponseDTO{
		UserID: adminID,
		Role:   "admin",
		Name:   name,
		Surname: surname,
		Email:  emailDB,
	}, nil
}

func (r *AuthRepository) GetUserByID(ctx context.Context, tx pgx.Tx, userID int, role string) (*dto.AuthResponseDTO, error) {
	var query string
	var name, surname, email sql.NullString

	switch role {
	case "student":
		query = `SELECT name, surname, email FROM student WHERE student_id = $1`
	case "teacher":
		query = `SELECT name, surname, email FROM teacher WHERE teacher_id = $1`
	case "admin":
		query = `SELECT name, surname, email FROM admin WHERE admin_id = $1`
	default:
		return nil, errors.New("invalid role")
	}

	err := tx.QueryRow(ctx, query, userID).Scan(&name, &surname, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &dto.AuthResponseDTO{
		UserID: userID,
		Role:   role,
		Name:   name.String,
		Surname: surname.String,
		Email:  email.String,
	}, nil
}


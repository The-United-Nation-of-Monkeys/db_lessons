package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// TransactionsCoursesRepository
type TransactionsCoursesRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, transactionID, courseID int) (*dto.TransactionsCoursesDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TransactionsCoursesDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, transactionID, courseID int) error
}

type TransactionsCoursesRepository struct{}

func NewTransactionsCoursesRepository() *TransactionsCoursesRepository {
	return &TransactionsCoursesRepository{}
}

func (r *TransactionsCoursesRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error) {
	query := `INSERT INTO transactions_courses (transaction_id, course_id) VALUES ($1, $2) RETURNING transaction_id, course_id`
	var relation dto.TransactionsCoursesDTO
	err := conn.QueryRow(ctx, query, data.TransactionID, data.CourseID).Scan(&relation.TransactionID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TransactionsCoursesRepository) GetByID(ctx context.Context, conn *pgx.Conn, transactionID, courseID int) (*dto.TransactionsCoursesDTO, error) {
	query := `SELECT transaction_id, course_id FROM transactions_courses WHERE transaction_id = $1 AND course_id = $2`
	var relation dto.TransactionsCoursesDTO
	err := conn.QueryRow(ctx, query, transactionID, courseID).Scan(&relation.TransactionID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TransactionsCoursesRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TransactionsCoursesDTO, error) {
	query := `SELECT transaction_id, course_id FROM transactions_courses`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.TransactionsCoursesDTO
	for rows.Next() {
		var relation dto.TransactionsCoursesDTO
		if err := rows.Scan(&relation.TransactionID, &relation.CourseID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *TransactionsCoursesRepository) Delete(ctx context.Context, conn *pgx.Conn, transactionID, courseID int) error {
	query := `DELETE FROM transactions_courses WHERE transaction_id = $1 AND course_id = $2`
	_, err := conn.Exec(ctx, query, transactionID, courseID)
	return err
}

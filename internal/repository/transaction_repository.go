package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.TransactionDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TransactionDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
	query := `INSERT INTO transaction (student_id, status_id, total_price) 
			  VALUES ($1, $2, $3) 
			  RETURNING transaction_id, student_id, status_id, total_price`
	var transaction dto.TransactionDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.StatusID, data.TotalPrice).
		Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.TransactionDTO, error) {
	query := `SELECT transaction_id, student_id, status_id, total_price FROM transaction WHERE transaction_id = $1`
	var transaction dto.TransactionDTO
	err := tx.QueryRow(ctx, query, id).Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TransactionDTO, error) {
	query := `SELECT transaction_id, student_id, status_id, total_price FROM transaction`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*dto.TransactionDTO
	for rows.Next() {
		var transaction dto.TransactionDTO
		if err := rows.Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice); err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (r *TransactionRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error) {
	query := `UPDATE transaction SET 
			  student_id = COALESCE($1, student_id),
			  status_id = COALESCE($2, status_id),
			  total_price = COALESCE($3, total_price)
			  WHERE transaction_id = $4
			  RETURNING transaction_id, student_id, status_id, total_price`
	var transaction dto.TransactionDTO
	err := tx.QueryRow(ctx, query, data.StudentID, data.StatusID, data.TotalPrice, id).
		Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM transaction WHERE transaction_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

func (r *TransactionRepository) GetReportByParams(ctx context.Context, tx pgx.Tx, params *dto.TransactionReportRequestDTO) ([]*dto.TransactionReportDTO, error) {
	query := `SELECT *
			  FROM get_transactions_report_dynamic($1, $2, $3)`
	response := make([]*dto.TransactionReportDTO, 0)
	rows, err := tx.Query(ctx, query, params.StatusName, params.MinTotal, params.MaxTotal)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var report dto.TransactionReportDTO
		if err = rows.Scan(report.TransactionsID, report.StudentID, report.StudentName, report.StudentSurname, report.StatusName, report.TotalPrice, report.CoursesNames); err != nil {
			return nil, err
		}
		response = append(response, &report)
	}

	return response, nil
}

func (r *TransactionRepository) GetReport(ctx context.Context, tx pgx.Tx) ([]*dto.TransactionReportDTO, error) {

}

package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.TransactionDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TransactionDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
	GetReportByParams(ctx context.Context, conn *pgx.Conn, params *dto.TransactionReportRequestDTO) ([]*dto.TransactionReportDTO, error)
	BulkUpdateTransactionStatus(ctx context.Context, conn *pgx.Conn, params *dto.BulkUpdateTransactionStatusDTO) error
}

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateTransactionDTO) (*dto.TransactionDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO transaction (student_id, status_id, total_price) 
			  VALUES ($1, $2, $3) 
			  RETURNING transaction_id, student_id, status_id, total_price`
	var transaction dto.TransactionDTO
	err = tx.QueryRow(ctx, query, data.StudentID, data.StatusID, data.TotalPrice).
		Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.TransactionDTO, error) {
	query := `SELECT transaction_id, student_id, status_id, total_price FROM transaction WHERE transaction_id = $1`
	var transaction dto.TransactionDTO
	err := conn.QueryRow(ctx, query, id).Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.TransactionDTO, error) {
	query := `SELECT transaction_id, student_id, status_id, total_price FROM transaction`
	rows, err := conn.Query(ctx, query)
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

func (r *TransactionRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateTransactionDTO) (*dto.TransactionDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE transaction SET 
			  student_id = COALESCE($1, student_id),
			  status_id = COALESCE($2, status_id),
			  total_price = COALESCE($3, total_price)
			  WHERE transaction_id = $4
			  RETURNING transaction_id, student_id, status_id, total_price`
	var transaction dto.TransactionDTO
	err = tx.QueryRow(ctx, query, data.StudentID, data.StatusID, data.TotalPrice, id).
		Scan(&transaction.TransactionID, &transaction.StudentID, &transaction.StatusID, &transaction.TotalPrice)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM transaction WHERE transaction_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *TransactionRepository) GetReportByParams(ctx context.Context, conn *pgx.Conn, params *dto.TransactionReportRequestDTO) ([]*dto.TransactionReportDTO, error) {
	query := `SELECT 
		transaction_id,
		student_id,
		student_name,
		student_surname,
		status_name,
		total_price,
		courses_names
	FROM get_transactions_report_dynamic($1, $2, $3)`

	response := make([]*dto.TransactionReportDTO, 0)

	var statusName *string
	if params.StatusName != "" {
		statusName = &params.StatusName
	}
	var minTotal *int
	if params.MinTotal > 0 {
		minTotal = &params.MinTotal
	}
	var maxTotal *int
	if params.MaxTotal > 0 {
		maxTotal = &params.MaxTotal
	}

	rows, err := conn.Query(ctx, query, statusName, minTotal, maxTotal)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var report dto.TransactionReportDTO
		var studentName, studentSurname, statusName, coursesNames *string

		err = rows.Scan(
			&report.TransactionsID,
			&report.StudentID,
			&studentName,
			&studentSurname,
			&statusName,
			&report.TotalPrice,
			&coursesNames,
		)
		if err != nil {
			return nil, err
		}

		report.StudentName = studentName
		report.StudentSurname = studentSurname
		report.StatusName = statusName
		report.CoursesNames = coursesNames

		response = append(response, &report)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *TransactionRepository) BulkUpdateTransactionStatus(ctx context.Context, conn *pgx.Conn, params *dto.BulkUpdateTransactionStatusDTO) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `CALL bulk_update_transaction_status($1, $2, $3, $4)`

	var minTotal *int
	if params.MinTotal > 0 {
		minTotal = &params.MinTotal
	}
	var maxTotal *int
	if params.MaxTotal > 0 {
		maxTotal = &params.MaxTotal
	}

	_, err = tx.Exec(ctx, query, params.OldStatusID, params.NewStatusID, minTotal, maxTotal)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

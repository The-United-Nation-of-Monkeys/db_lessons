package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type StatusHomeworkRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusHomeworkDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusHomeworkDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type StatusHomeworkRepository struct{}

func NewStatusHomeworkRepository() *StatusHomeworkRepository {
	return &StatusHomeworkRepository{}
}

func (r *StatusHomeworkRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	query := `INSERT INTO status_homework (name) VALUES ($1) RETURNING status_homework_id, name`
	var status dto.StatusHomeworkDTO
	err := conn.QueryRow(ctx, query, data.Name).Scan(&status.StatusHomeworkID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusHomeworkRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusHomeworkDTO, error) {
	query := `SELECT status_homework_id, name FROM status_homework WHERE status_homework_id = $1`
	var status dto.StatusHomeworkDTO
	err := conn.QueryRow(ctx, query, id).Scan(&status.StatusHomeworkID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusHomeworkRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusHomeworkDTO, error) {
	query := `SELECT status_homework_id, name FROM status_homework`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*dto.StatusHomeworkDTO
	for rows.Next() {
		var status dto.StatusHomeworkDTO
		if err := rows.Scan(&status.StatusHomeworkID, &status.Name); err != nil {
			return nil, err
		}
		statuses = append(statuses, &status)
	}
	return statuses, nil
}

func (r *StatusHomeworkRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusHomeworkDTO) (*dto.StatusHomeworkDTO, error) {
	query := `UPDATE status_homework SET name = COALESCE($1, name) WHERE status_homework_id = $2 RETURNING status_homework_id, name`
	var status dto.StatusHomeworkDTO
	err := conn.QueryRow(ctx, query, data.Name, id).Scan(&status.StatusHomeworkID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusHomeworkRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	query := `DELETE FROM status_homework WHERE status_homework_id = $1`
	_, err := conn.Exec(ctx, query, id)
	return err
}

type StatusAnswerRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusAnswerDTO) (*dto.StatusAnswerDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusAnswerDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusAnswerDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusAnswerDTO) (*dto.StatusAnswerDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type StatusAnswerRepository struct{}

func NewStatusAnswerRepository() *StatusAnswerRepository {
	return &StatusAnswerRepository{}
}

func (r *StatusAnswerRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusAnswerDTO) (*dto.StatusAnswerDTO, error) {
	query := `INSERT INTO status_answer (name) VALUES ($1) RETURNING status_answer_id, name`
	var status dto.StatusAnswerDTO
	err := conn.QueryRow(ctx, query, data.Name).Scan(&status.StatusAnswerID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusAnswerRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusAnswerDTO, error) {
	query := `SELECT status_answer_id, name FROM status_answer WHERE status_answer_id = $1`
	var status dto.StatusAnswerDTO
	err := conn.QueryRow(ctx, query, id).Scan(&status.StatusAnswerID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusAnswerRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusAnswerDTO, error) {
	query := `SELECT status_answer_id, name FROM status_answer`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*dto.StatusAnswerDTO
	for rows.Next() {
		var status dto.StatusAnswerDTO
		if err := rows.Scan(&status.StatusAnswerID, &status.Name); err != nil {
			return nil, err
		}
		statuses = append(statuses, &status)
	}
	return statuses, nil
}

func (r *StatusAnswerRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusAnswerDTO) (*dto.StatusAnswerDTO, error) {
	query := `UPDATE status_answer SET name = COALESCE($1, name) WHERE status_answer_id = $2 RETURNING status_answer_id, name`
	var status dto.StatusAnswerDTO
	err := conn.QueryRow(ctx, query, data.Name, id).Scan(&status.StatusAnswerID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusAnswerRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	query := `DELETE FROM status_answer WHERE status_answer_id = $1`
	_, err := conn.Exec(ctx, query, id)
	return err
}

type StatusTransactionRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusTransactionDTO) (*dto.StatusTransactionDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusTransactionDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusTransactionDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusTransactionDTO) (*dto.StatusTransactionDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type StatusTransactionRepository struct{}

func NewStatusTransactionRepository() *StatusTransactionRepository {
	return &StatusTransactionRepository{}
}

func (r *StatusTransactionRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateStatusTransactionDTO) (*dto.StatusTransactionDTO, error) {
	query := `INSERT INTO status_transaction (name) VALUES ($1) RETURNING status_transaction_id, name`
	var status dto.StatusTransactionDTO
	err := conn.QueryRow(ctx, query, data.Name).Scan(&status.StatusTransactionID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusTransactionRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.StatusTransactionDTO, error) {
	query := `SELECT status_transaction_id, name FROM status_transaction WHERE status_transaction_id = $1`
	var status dto.StatusTransactionDTO
	err := conn.QueryRow(ctx, query, id).Scan(&status.StatusTransactionID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusTransactionRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StatusTransactionDTO, error) {
	query := `SELECT status_transaction_id, name FROM status_transaction`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*dto.StatusTransactionDTO
	for rows.Next() {
		var status dto.StatusTransactionDTO
		if err := rows.Scan(&status.StatusTransactionID, &status.Name); err != nil {
			return nil, err
		}
		statuses = append(statuses, &status)
	}
	return statuses, nil
}

func (r *StatusTransactionRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateStatusTransactionDTO) (*dto.StatusTransactionDTO, error) {
	query := `UPDATE status_transaction SET name = COALESCE($1, name) WHERE status_transaction_id = $2 RETURNING status_transaction_id, name`
	var status dto.StatusTransactionDTO
	err := conn.QueryRow(ctx, query, data.Name, id).Scan(&status.StatusTransactionID, &status.Name)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *StatusTransactionRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	query := `DELETE FROM status_transaction WHERE status_transaction_id = $1`
	_, err := conn.Exec(ctx, query, id)
	return err
}

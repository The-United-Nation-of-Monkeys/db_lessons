package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type CurrencyRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CurrencyDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CurrencyDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCurrencyDTO) (*dto.CurrencyDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type CurrencyRepository struct{}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (r *CurrencyRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCurrencyDTO) (*dto.CurrencyDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO currency (name) VALUES ($1) RETURNING currency_id, name`
	var currency dto.CurrencyDTO
	err = tx.QueryRow(ctx, query, data.Name).Scan(&currency.CurrencyID, &currency.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &currency, nil
}

func (r *CurrencyRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CurrencyDTO, error) {
	query := `SELECT currency_id, name FROM currency WHERE currency_id = $1`
	var currency dto.CurrencyDTO
	err := conn.QueryRow(ctx, query, id).Scan(&currency.CurrencyID, &currency.Name)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *CurrencyRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CurrencyDTO, error) {
	query := `SELECT currency_id, name FROM currency`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []*dto.CurrencyDTO
	for rows.Next() {
		var currency dto.CurrencyDTO
		if err := rows.Scan(&currency.CurrencyID, &currency.Name); err != nil {
			return nil, err
		}
		currencies = append(currencies, &currency)
	}
	return currencies, nil
}

func (r *CurrencyRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCurrencyDTO) (*dto.CurrencyDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE currency SET name = COALESCE($1, name) WHERE currency_id = $2 RETURNING currency_id, name`
	var currency dto.CurrencyDTO
	err = tx.QueryRow(ctx, query, data.Name, id).Scan(&currency.CurrencyID, &currency.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &currency, nil
}

func (r *CurrencyRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM currency WHERE currency_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

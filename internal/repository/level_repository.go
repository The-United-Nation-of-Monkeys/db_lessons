package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type LevelRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLevelDTO) (*dto.LevelDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.LevelDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LevelDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type LevelRepository struct{}

func NewLevelRepository() *LevelRepository {
	return &LevelRepository{}
}

func (r *LevelRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLevelDTO) (*dto.LevelDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO level (name) VALUES ($1) RETURNING level_id, name`
	var level dto.LevelDTO
	err = tx.QueryRow(ctx, query, data.Name).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &level, nil
}

func (r *LevelRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.LevelDTO, error) {
	query := `SELECT level_id, name FROM level WHERE level_id = $1`
	var level dto.LevelDTO
	err := conn.QueryRow(ctx, query, id).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

func (r *LevelRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LevelDTO, error) {
	query := `SELECT level_id, name FROM level`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var levels []*dto.LevelDTO
	for rows.Next() {
		var level dto.LevelDTO
		if err := rows.Scan(&level.LevelID, &level.Name); err != nil {
			return nil, err
		}
		levels = append(levels, &level)
	}
	return levels, nil
}

func (r *LevelRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE level SET name = COALESCE($1, name) WHERE level_id = $2 RETURNING level_id, name`
	var level dto.LevelDTO
	err = tx.QueryRow(ctx, query, data.Name, id).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &level, nil
}

func (r *LevelRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM level WHERE level_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

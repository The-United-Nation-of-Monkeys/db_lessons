package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type LevelRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLevelDTO) (*dto.LevelDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.LevelDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LevelDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type LevelRepository struct{}

func NewLevelRepository() *LevelRepository {
	return &LevelRepository{}
}

func (r *LevelRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLevelDTO) (*dto.LevelDTO, error) {
	query := `INSERT INTO level (name) VALUES ($1) RETURNING level_id, name`
	var level dto.LevelDTO
	err := tx.QueryRow(ctx, query, data.Name).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

func (r *LevelRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.LevelDTO, error) {
	query := `SELECT level_id, name FROM level WHERE level_id = $1`
	var level dto.LevelDTO
	err := tx.QueryRow(ctx, query, id).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

func (r *LevelRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LevelDTO, error) {
	query := `SELECT level_id, name FROM level`
	rows, err := tx.Query(ctx, query)
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

func (r *LevelRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateLevelDTO) (*dto.LevelDTO, error) {
	query := `UPDATE level SET name = COALESCE($1, name) WHERE level_id = $2 RETURNING level_id, name`
	var level dto.LevelDTO
	err := tx.QueryRow(ctx, query, data.Name, id).Scan(&level.LevelID, &level.Name)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

func (r *LevelRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM level WHERE level_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

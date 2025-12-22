package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type MaterialRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateMaterialDTO) (*dto.MaterialDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.MaterialDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.MaterialDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateMaterialDTO) (*dto.MaterialDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type MaterialRepository struct{}

func NewMaterialRepository() *MaterialRepository {
	return &MaterialRepository{}
}

func (r *MaterialRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateMaterialDTO) (*dto.MaterialDTO, error) {
	query := `INSERT INTO material (source, extension, size) 
			  VALUES ($1, $2, $3) 
			  RETURNING material_id, source, extension, size`
	var material dto.MaterialDTO
	err := tx.QueryRow(ctx, query, data.Source, data.Extension, data.Size).
		Scan(&material.MaterialID, &material.Source, &material.Extension, &material.Size)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *MaterialRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.MaterialDTO, error) {
	query := `SELECT material_id, source, extension, size FROM material WHERE material_id = $1`
	var material dto.MaterialDTO
	err := tx.QueryRow(ctx, query, id).Scan(&material.MaterialID, &material.Source, &material.Extension, &material.Size)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *MaterialRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.MaterialDTO, error) {
	query := `SELECT material_id, source, extension, size FROM material`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []*dto.MaterialDTO
	for rows.Next() {
		var material dto.MaterialDTO
		if err := rows.Scan(&material.MaterialID, &material.Source, &material.Extension, &material.Size); err != nil {
			return nil, err
		}
		materials = append(materials, &material)
	}
	return materials, nil
}

func (r *MaterialRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateMaterialDTO) (*dto.MaterialDTO, error) {
	query := `UPDATE material SET 
			  source = COALESCE($1, source),
			  extension = COALESCE($2, extension),
			  size = COALESCE($3, size)
			  WHERE material_id = $4
			  RETURNING material_id, source, extension, size`
	var material dto.MaterialDTO
	err := tx.QueryRow(ctx, query, data.Source, data.Extension, data.Size, id).
		Scan(&material.MaterialID, &material.Source, &material.Extension, &material.Size)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *MaterialRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM material WHERE material_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

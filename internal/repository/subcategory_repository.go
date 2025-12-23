package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type SubcategoryRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.SubcategoryDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.SubcategoryDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateSubcategoryDTO) (*dto.SubcategoryDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type SubcategoryRepository struct{}

func NewSubcategoryRepository() *SubcategoryRepository {
	return &SubcategoryRepository{}
}

func (r *SubcategoryRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateSubcategoryDTO) (*dto.SubcategoryDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO subcategory (name, category_id) VALUES ($1, $2) RETURNING subcategory_id, name, category_id`
	var subcategory dto.SubcategoryDTO
	err = tx.QueryRow(ctx, query, data.Name, data.CategoryID).Scan(&subcategory.SubcategoryID, &subcategory.Name, &subcategory.CategoryID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &subcategory, nil
}

func (r *SubcategoryRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.SubcategoryDTO, error) {
	query := `SELECT subcategory_id, name, category_id FROM subcategory WHERE subcategory_id = $1`
	var subcategory dto.SubcategoryDTO
	err := conn.QueryRow(ctx, query, id).Scan(&subcategory.SubcategoryID, &subcategory.Name, &subcategory.CategoryID)
	if err != nil {
		return nil, err
	}
	return &subcategory, nil
}

func (r *SubcategoryRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.SubcategoryDTO, error) {
	query := `SELECT subcategory_id, name, category_id FROM subcategory`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategories []*dto.SubcategoryDTO
	for rows.Next() {
		var subcategory dto.SubcategoryDTO
		if err := rows.Scan(&subcategory.SubcategoryID, &subcategory.Name, &subcategory.CategoryID); err != nil {
			return nil, err
		}
		subcategories = append(subcategories, &subcategory)
	}
	return subcategories, nil
}

func (r *SubcategoryRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateSubcategoryDTO) (*dto.SubcategoryDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE subcategory SET name = COALESCE($1, name), category_id = COALESCE($2, category_id) WHERE subcategory_id = $3 RETURNING subcategory_id, name, category_id`
	var subcategory dto.SubcategoryDTO
	err = tx.QueryRow(ctx, query, data.Name, data.CategoryID, id).Scan(&subcategory.SubcategoryID, &subcategory.Name, &subcategory.CategoryID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &subcategory, nil
}

func (r *SubcategoryRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM subcategory WHERE subcategory_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

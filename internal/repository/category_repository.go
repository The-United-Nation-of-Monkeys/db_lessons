package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type CategoryRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CategoryDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CategoryDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateCategoryDTO) (*dto.CategoryDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO category (name) VALUES ($1) RETURNING category_id, name`
	var category dto.CategoryDTO
	err = tx.QueryRow(ctx, query, data.Name).Scan(&category.CategoryID, &category.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.CategoryDTO, error) {
	query := `SELECT category_id, name FROM category WHERE category_id = $1`
	var category dto.CategoryDTO
	err := conn.QueryRow(ctx, query, id).Scan(&category.CategoryID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.CategoryDTO, error) {
	query := `SELECT category_id, name FROM category`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*dto.CategoryDTO
	for rows.Next() {
		var category dto.CategoryDTO
		if err := rows.Scan(&category.CategoryID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (r *CategoryRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateCategoryDTO) (*dto.CategoryDTO, error) {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE category SET name = COALESCE($1, name) WHERE category_id = $2 RETURNING category_id, name`
	var category dto.CategoryDTO
	err = tx.QueryRow(ctx, query, data.Name, id).Scan(&category.CategoryID, &category.Name)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	// Begin transaction for write operation
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM category WHERE category_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

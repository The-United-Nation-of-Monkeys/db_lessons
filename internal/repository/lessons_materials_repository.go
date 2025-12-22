package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// LessonsMaterialsRepository
type LessonsMaterialsRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonsMaterialsDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, lessonID, materialID int) error
}

type LessonsMaterialsRepository struct{}

func NewLessonsMaterialsRepository() *LessonsMaterialsRepository {
	return &LessonsMaterialsRepository{}
}

func (r *LessonsMaterialsRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error) {
	query := `INSERT INTO lessons_materials (lesson_id, material_id) VALUES ($1, $2) RETURNING lesson_id, material_id`
	var relation dto.LessonsMaterialsDTO
	err := tx.QueryRow(ctx, query, data.LessonID, data.MaterialID).Scan(&relation.LessonID, &relation.MaterialID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonsMaterialsRepository) GetByID(ctx context.Context, tx pgx.Tx, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error) {
	query := `SELECT lesson_id, material_id FROM lessons_materials WHERE lesson_id = $1 AND material_id = $2`
	var relation dto.LessonsMaterialsDTO
	err := tx.QueryRow(ctx, query, lessonID, materialID).Scan(&relation.LessonID, &relation.MaterialID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonsMaterialsRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonsMaterialsDTO, error) {
	query := `SELECT lesson_id, material_id FROM lessons_materials`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.LessonsMaterialsDTO
	for rows.Next() {
		var relation dto.LessonsMaterialsDTO
		if err := rows.Scan(&relation.LessonID, &relation.MaterialID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *LessonsMaterialsRepository) Delete(ctx context.Context, tx pgx.Tx, lessonID, materialID int) error {
	query := `DELETE FROM lessons_materials WHERE lesson_id = $1 AND material_id = $2`
	_, err := tx.Exec(ctx, query, lessonID, materialID)
	return err
}


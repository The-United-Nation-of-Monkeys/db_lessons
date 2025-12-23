package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type LessonsMaterialsRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LessonsMaterialsDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, lessonID, materialID int) error
}

type LessonsMaterialsRepository struct{}

func NewLessonsMaterialsRepository() *LessonsMaterialsRepository {
	return &LessonsMaterialsRepository{}
}

func (r *LessonsMaterialsRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error) {
	query := `INSERT INTO lessons_materials (lesson_id, material_id) VALUES ($1, $2) RETURNING lesson_id, material_id`
	var relation dto.LessonsMaterialsDTO
	err := conn.QueryRow(ctx, query, data.LessonID, data.MaterialID).Scan(&relation.LessonID, &relation.MaterialID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonsMaterialsRepository) GetByID(ctx context.Context, conn *pgx.Conn, lessonID, materialID int) (*dto.LessonsMaterialsDTO, error) {
	query := `SELECT lesson_id, material_id FROM lessons_materials WHERE lesson_id = $1 AND material_id = $2`
	var relation dto.LessonsMaterialsDTO
	err := conn.QueryRow(ctx, query, lessonID, materialID).Scan(&relation.LessonID, &relation.MaterialID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonsMaterialsRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LessonsMaterialsDTO, error) {
	query := `SELECT lesson_id, material_id FROM lessons_materials`
	rows, err := conn.Query(ctx, query)
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

func (r *LessonsMaterialsRepository) Delete(ctx context.Context, conn *pgx.Conn, lessonID, materialID int) error {
	query := `DELETE FROM lessons_materials WHERE lesson_id = $1 AND material_id = $2`
	_, err := conn.Exec(ctx, query, lessonID, materialID)
	return err
}


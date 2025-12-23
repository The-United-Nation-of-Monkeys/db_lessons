package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type LessonRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLessonDTO) (*dto.LessonDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.LessonDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LessonDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateLessonDTO) (*dto.LessonDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type LessonRepository struct{}

func NewLessonRepository() *LessonRepository {
	return &LessonRepository{}
}

func (r *LessonRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateLessonDTO) (*dto.LessonDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO lesson (name, description, open_time, video_link, lesson_text) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING lesson_id, name, description, open_time, video_link, lesson_text`
	var lesson dto.LessonDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.OpenTime, data.VideoLink, data.LessonText).
		Scan(&lesson.LessonID, &lesson.Name, &lesson.Description, &lesson.OpenTime, &lesson.VideoLink, &lesson.LessonText)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *LessonRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.LessonDTO, error) {
	query := `SELECT lesson_id, name, description, open_time, video_link, lesson_text FROM lesson WHERE lesson_id = $1`
	var lesson dto.LessonDTO
	err := conn.QueryRow(ctx, query, id).Scan(&lesson.LessonID, &lesson.Name, &lesson.Description, &lesson.OpenTime, &lesson.VideoLink, &lesson.LessonText)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *LessonRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.LessonDTO, error) {
	query := `SELECT lesson_id, name, description, open_time, video_link, lesson_text FROM lesson`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*dto.LessonDTO
	for rows.Next() {
		var lesson dto.LessonDTO
		if err := rows.Scan(&lesson.LessonID, &lesson.Name, &lesson.Description, &lesson.OpenTime, &lesson.VideoLink, &lesson.LessonText); err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}
	return lessons, nil
}

func (r *LessonRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateLessonDTO) (*dto.LessonDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE lesson SET 
			  name = COALESCE($1, name),
			  description = COALESCE($2, description),
			  open_time = COALESCE($3, open_time),
			  video_link = COALESCE($4, video_link),
			  lesson_text = COALESCE($5, lesson_text)
			  WHERE lesson_id = $6
			  RETURNING lesson_id, name, description, open_time, video_link, lesson_text`
	var lesson dto.LessonDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.OpenTime, data.VideoLink, data.LessonText, id).
		Scan(&lesson.LessonID, &lesson.Name, &lesson.Description, &lesson.OpenTime, &lesson.VideoLink, &lesson.LessonText)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *LessonRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM lesson WHERE lesson_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

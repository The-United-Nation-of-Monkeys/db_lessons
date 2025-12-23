package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type HomeworkRepositoryInterface interface {
	Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error)
	GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.HomeworkDTO, error)
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.HomeworkDTO, error)
	Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error)
	Delete(ctx context.Context, conn *pgx.Conn, id int) error
}

type HomeworkRepository struct{}

func NewHomeworkRepository() *HomeworkRepository {
	return &HomeworkRepository{}
}

func (r *HomeworkRepository) Create(ctx context.Context, conn *pgx.Conn, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO homework (name, description, deadline_time) 
			  VALUES ($1, $2, $3) 
			  RETURNING homework_id, name, description, deadline_time`
	var homework dto.HomeworkDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.DeadlineTime).
		Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &homework, nil
}

func (r *HomeworkRepository) GetByID(ctx context.Context, conn *pgx.Conn, id int) (*dto.HomeworkDTO, error) {
	query := `SELECT homework_id, name, description, deadline_time FROM homework WHERE homework_id = $1`
	var homework dto.HomeworkDTO
	err := conn.QueryRow(ctx, query, id).Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}
	return &homework, nil
}

func (r *HomeworkRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.HomeworkDTO, error) {
	query := `SELECT homework_id, name, description, deadline_time FROM homework`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var homeworks []*dto.HomeworkDTO
	for rows.Next() {
		var homework dto.HomeworkDTO
		if err := rows.Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime); err != nil {
			return nil, err
		}
		homeworks = append(homeworks, &homework)
	}
	return homeworks, nil
}

func (r *HomeworkRepository) Update(ctx context.Context, conn *pgx.Conn, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE homework SET 
			  name = COALESCE($1, name),
			  description = COALESCE($2, description),
			  deadline_time = COALESCE($3, deadline_time)
			  WHERE homework_id = $4
			  RETURNING homework_id, name, description, deadline_time`
	var homework dto.HomeworkDTO
	err = tx.QueryRow(ctx, query, data.Name, data.Description, data.DeadlineTime, id).
		Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &homework, nil
}

func (r *HomeworkRepository) Delete(ctx context.Context, conn *pgx.Conn, id int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM homework WHERE homework_id = $1`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

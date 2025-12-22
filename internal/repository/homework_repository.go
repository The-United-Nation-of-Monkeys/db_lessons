package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type HomeworkRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.HomeworkDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworkDTO, error)
	Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id int) error
}

type HomeworkRepository struct{}

func NewHomeworkRepository() *HomeworkRepository {
	return &HomeworkRepository{}
}

func (r *HomeworkRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworkDTO) (*dto.HomeworkDTO, error) {
	query := `INSERT INTO homework (name, description, deadline_time) 
			  VALUES ($1, $2, $3) 
			  RETURNING homework_id, name, description, deadline_time`
	var homework dto.HomeworkDTO
	err := tx.QueryRow(ctx, query, data.Name, data.Description, data.DeadlineTime).
		Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}
	return &homework, nil
}

func (r *HomeworkRepository) GetByID(ctx context.Context, tx pgx.Tx, id int) (*dto.HomeworkDTO, error) {
	query := `SELECT homework_id, name, description, deadline_time FROM homework WHERE homework_id = $1`
	var homework dto.HomeworkDTO
	err := tx.QueryRow(ctx, query, id).Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}
	return &homework, nil
}

func (r *HomeworkRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworkDTO, error) {
	query := `SELECT homework_id, name, description, deadline_time FROM homework`
	rows, err := tx.Query(ctx, query)
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

func (r *HomeworkRepository) Update(ctx context.Context, tx pgx.Tx, id int, data *dto.UpdateHomeworkDTO) (*dto.HomeworkDTO, error) {
	query := `UPDATE homework SET 
			  name = COALESCE($1, name),
			  description = COALESCE($2, description),
			  deadline_time = COALESCE($3, deadline_time)
			  WHERE homework_id = $4
			  RETURNING homework_id, name, description, deadline_time`
	var homework dto.HomeworkDTO
	err := tx.QueryRow(ctx, query, data.Name, data.Description, data.DeadlineTime, id).
		Scan(&homework.HomeworkID, &homework.Name, &homework.Description, &homework.DeadlineTime)
	if err != nil {
		return nil, err
	}
	return &homework, nil
}

func (r *HomeworkRepository) Delete(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE FROM homework WHERE homework_id = $1`
	_, err := tx.Exec(ctx, query, id)
	return err
}

package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type StudentCategoryStatsRepositoryInterface interface {
	GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StudentCategoryStatsDTO, error)
	GetByStudentID(ctx context.Context, conn *pgx.Conn, studentID int) ([]*dto.StudentCategoryStatsDTO, error)
	GetByCategoryID(ctx context.Context, conn *pgx.Conn, categoryID int) ([]*dto.StudentCategoryStatsDTO, error)
}

type StudentCategoryStatsRepository struct{}

func NewStudentCategoryStatsRepository() *StudentCategoryStatsRepository {
	return &StudentCategoryStatsRepository{}
}

func (r *StudentCategoryStatsRepository) GetAll(ctx context.Context, conn *pgx.Conn) ([]*dto.StudentCategoryStatsDTO, error) {
	query := `SELECT 
		student_id,
		student_name,
		student_surname,
		category_id,
		category_name,
		solved_tasks_count,
		points_earned,
		points_possible,
		success_percent
	FROM v_student_category_stats`
	
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*dto.StudentCategoryStatsDTO
	for rows.Next() {
		var stat dto.StudentCategoryStatsDTO
		err = rows.Scan(
			&stat.StudentID,
			&stat.StudentName,
			&stat.StudentSurname,
			&stat.CategoryID,
			&stat.CategoryName,
			&stat.SolvedTasksCount,
			&stat.PointsEarned,
			&stat.PointsPossible,
			&stat.SuccessPercent,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, &stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *StudentCategoryStatsRepository) GetByStudentID(ctx context.Context, conn *pgx.Conn, studentID int) ([]*dto.StudentCategoryStatsDTO, error) {
	query := `SELECT 
		student_id,
		student_name,
		student_surname,
		category_id,
		category_name,
		solved_tasks_count,
		points_earned,
		points_possible,
		success_percent
	FROM v_student_category_stats
	WHERE student_id = $1`
	
	rows, err := conn.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*dto.StudentCategoryStatsDTO
	for rows.Next() {
		var stat dto.StudentCategoryStatsDTO
		err = rows.Scan(
			&stat.StudentID,
			&stat.StudentName,
			&stat.StudentSurname,
			&stat.CategoryID,
			&stat.CategoryName,
			&stat.SolvedTasksCount,
			&stat.PointsEarned,
			&stat.PointsPossible,
			&stat.SuccessPercent,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, &stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *StudentCategoryStatsRepository) GetByCategoryID(ctx context.Context, conn *pgx.Conn, categoryID int) ([]*dto.StudentCategoryStatsDTO, error) {
	query := `SELECT 
		student_id,
		student_name,
		student_surname,
		category_id,
		category_name,
		solved_tasks_count,
		points_earned,
		points_possible,
		success_percent
	FROM v_student_category_stats
	WHERE category_id = $1`
	
	rows, err := conn.Query(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*dto.StudentCategoryStatsDTO
	for rows.Next() {
		var stat dto.StudentCategoryStatsDTO
		err = rows.Scan(
			&stat.StudentID,
			&stat.StudentName,
			&stat.StudentSurname,
			&stat.CategoryID,
			&stat.CategoryName,
			&stat.SolvedTasksCount,
			&stat.PointsEarned,
			&stat.PointsPossible,
			&stat.SuccessPercent,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, &stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}


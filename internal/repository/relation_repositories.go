package repository

import (
	"context"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

// TeachersCoursesRepository
type TeachersCoursesRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeachersCoursesDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, courseID, teacherID int) error
}

type TeachersCoursesRepository struct{}

func NewTeachersCoursesRepository() *TeachersCoursesRepository {
	return &TeachersCoursesRepository{}
}

func (r *TeachersCoursesRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTeachersCoursesDTO) (*dto.TeachersCoursesDTO, error) {
	query := `INSERT INTO teachers_courses (course_id, teacher_id) VALUES ($1, $2) RETURNING course_id, teacher_id`
	var relation dto.TeachersCoursesDTO
	err := tx.QueryRow(ctx, query, data.CourseID, data.TeacherID).Scan(&relation.CourseID, &relation.TeacherID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TeachersCoursesRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TeachersCoursesDTO, error) {
	query := `SELECT course_id, teacher_id FROM teachers_courses`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.TeachersCoursesDTO
	for rows.Next() {
		var relation dto.TeachersCoursesDTO
		if err := rows.Scan(&relation.CourseID, &relation.TeacherID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *TeachersCoursesRepository) Delete(ctx context.Context, tx pgx.Tx, courseID, teacherID int) error {
	query := `DELETE FROM teachers_courses WHERE course_id = $1 AND teacher_id = $2`
	_, err := tx.Exec(ctx, query, courseID, teacherID)
	return err
}

// CourseLessonsRepository
type CourseLessonsRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.CourseLessonsDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, lessonID int) error
}

type CourseLessonsRepository struct{}

func NewCourseLessonsRepository() *CourseLessonsRepository {
	return &CourseLessonsRepository{}
}

func (r *CourseLessonsRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateCourseLessonsDTO) (*dto.CourseLessonsDTO, error) {
	query := `INSERT INTO course_lessons (lesson_id, course_id) VALUES ($1, $2) RETURNING lesson_id, course_id`
	var relation dto.CourseLessonsDTO
	err := tx.QueryRow(ctx, query, data.LessonID, data.CourseID).Scan(&relation.LessonID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *CourseLessonsRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.CourseLessonsDTO, error) {
	query := `SELECT lesson_id, course_id FROM course_lessons`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.CourseLessonsDTO
	for rows.Next() {
		var relation dto.CourseLessonsDTO
		if err := rows.Scan(&relation.LessonID, &relation.CourseID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *CourseLessonsRepository) Delete(ctx context.Context, tx pgx.Tx, lessonID int) error {
	query := `DELETE FROM course_lessons WHERE lesson_id = $1`
	_, err := tx.Exec(ctx, query, lessonID)
	return err
}

// LessonsMaterialsRepository
type LessonsMaterialsRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonsMaterialsDTO) (*dto.LessonsMaterialsDTO, error)
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

// LessonHomeworksRepository
type LessonHomeworksRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonHomeworksDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) error
}

type LessonHomeworksRepository struct{}

func NewLessonHomeworksRepository() *LessonHomeworksRepository {
	return &LessonHomeworksRepository{}
}

func (r *LessonHomeworksRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateLessonHomeworksDTO) (*dto.LessonHomeworksDTO, error) {
	query := `INSERT INTO lesson_homeworks (lesson_id, homework_id) VALUES ($1, $2) RETURNING lesson_id, homework_id`
	var relation dto.LessonHomeworksDTO
	err := tx.QueryRow(ctx, query, data.LessonID, data.HomeworkID).Scan(&relation.LessonID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *LessonHomeworksRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.LessonHomeworksDTO, error) {
	query := `SELECT lesson_id, homework_id FROM lesson_homeworks`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.LessonHomeworksDTO
	for rows.Next() {
		var relation dto.LessonHomeworksDTO
		if err := rows.Scan(&relation.LessonID, &relation.HomeworkID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *LessonHomeworksRepository) Delete(ctx context.Context, tx pgx.Tx, lessonID, homeworkID int) error {
	query := `DELETE FROM lesson_homeworks WHERE lesson_id = $1 AND homework_id = $2`
	_, err := tx.Exec(ctx, query, lessonID, homeworkID)
	return err
}

// HomeworksTasksRepository
type HomeworksTasksRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworksTasksDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) error
}

type HomeworksTasksRepository struct{}

func NewHomeworksTasksRepository() *HomeworksTasksRepository {
	return &HomeworksTasksRepository{}
}

func (r *HomeworksTasksRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateHomeworksTasksDTO) (*dto.HomeworksTasksDTO, error) {
	query := `INSERT INTO homeworks_tasks (task_id, homework_id) VALUES ($1, $2) RETURNING task_id, homework_id`
	var relation dto.HomeworksTasksDTO
	err := tx.QueryRow(ctx, query, data.TaskID, data.HomeworkID).Scan(&relation.TaskID, &relation.HomeworkID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *HomeworksTasksRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.HomeworksTasksDTO, error) {
	query := `SELECT task_id, homework_id FROM homeworks_tasks`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.HomeworksTasksDTO
	for rows.Next() {
		var relation dto.HomeworksTasksDTO
		if err := rows.Scan(&relation.TaskID, &relation.HomeworkID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *HomeworksTasksRepository) Delete(ctx context.Context, tx pgx.Tx, taskID, homeworkID int) error {
	query := `DELETE FROM homeworks_tasks WHERE task_id = $1 AND homework_id = $2`
	_, err := tx.Exec(ctx, query, taskID, homeworkID)
	return err
}

// TransactionsCoursesRepository
type TransactionsCoursesRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error)
	GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TransactionsCoursesDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, transactionID, courseID int) error
}

type TransactionsCoursesRepository struct{}

func NewTransactionsCoursesRepository() *TransactionsCoursesRepository {
	return &TransactionsCoursesRepository{}
}

func (r *TransactionsCoursesRepository) Create(ctx context.Context, tx pgx.Tx, data *dto.CreateTransactionsCoursesDTO) (*dto.TransactionsCoursesDTO, error) {
	query := `INSERT INTO transactions_courses (transaction_id, course_id) VALUES ($1, $2) RETURNING transaction_id, course_id`
	var relation dto.TransactionsCoursesDTO
	err := tx.QueryRow(ctx, query, data.TransactionID, data.CourseID).Scan(&relation.TransactionID, &relation.CourseID)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r *TransactionsCoursesRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]*dto.TransactionsCoursesDTO, error) {
	query := `SELECT transaction_id, course_id FROM transactions_courses`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relations []*dto.TransactionsCoursesDTO
	for rows.Next() {
		var relation dto.TransactionsCoursesDTO
		if err := rows.Scan(&relation.TransactionID, &relation.CourseID); err != nil {
			return nil, err
		}
		relations = append(relations, &relation)
	}
	return relations, nil
}

func (r *TransactionsCoursesRepository) Delete(ctx context.Context, tx pgx.Tx, transactionID, courseID int) error {
	query := `DELETE FROM transactions_courses WHERE transaction_id = $1 AND course_id = $2`
	_, err := tx.Exec(ctx, query, transactionID, courseID)
	return err
}

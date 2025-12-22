package initializer

import (
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceList struct {
	StudentService           service.StudentServiceInterface
	TeacherService           service.TeacherServiceInterface
	CategoryService          service.CategoryServiceInterface
	CurrencyService          service.CurrencyServiceInterface
	LevelService             service.LevelServiceInterface
	SubcategoryService       service.SubcategoryServiceInterface
	MaterialService          service.MaterialServiceInterface
	HomeworkService          service.HomeworkServiceInterface
	LessonService            service.LessonServiceInterface
	CourseService            service.CourseServiceInterface
	TaskService              service.TaskServiceInterface
	StatusHomeworkService    service.StatusHomeworkServiceInterface
	StatusAnswerService      service.StatusAnswerServiceInterface
	StatusTransactionService service.StatusTransactionServiceInterface
	StudentAnswerService     service.StudentAnswerServiceInterface
	HomeworkResultService    service.HomeworkResultServiceInterface
	TransactionService       service.TransactionServiceInterface
}

func NewServiceList(repositories *RepositoryList, dbPool *pgxpool.Pool) *ServiceList {
	return &ServiceList{
		StudentService:           service.NewStudentService(dbPool, repositories.StudentRepository),
		TeacherService:           service.NewTeacherService(dbPool, repositories.TeacherRepository),
		CategoryService:          service.NewCategoryService(dbPool, repositories.CategoryRepository),
		CurrencyService:          service.NewCurrencyService(dbPool, repositories.CurrencyRepository),
		LevelService:             service.NewLevelService(dbPool, repositories.LevelRepository),
		SubcategoryService:       service.NewSubcategoryService(dbPool, repositories.SubcategoryRepository),
		MaterialService:          service.NewMaterialService(dbPool, repositories.MaterialRepository),
		HomeworkService:          service.NewHomeworkService(dbPool, repositories.HomeworkRepository),
		LessonService:            service.NewLessonService(dbPool, repositories.LessonRepository),
		CourseService:            service.NewCourseService(dbPool, repositories.CourseRepository),
		TaskService:              service.NewTaskService(dbPool, repositories.TaskRepository),
		StatusHomeworkService:    service.NewStatusHomeworkService(dbPool, repositories.StatusHomeworkRepository),
		StatusAnswerService:      service.NewStatusAnswerService(dbPool, repositories.StatusAnswerRepository),
		StatusTransactionService: service.NewStatusTransactionService(dbPool, repositories.StatusTransactionRepository),
		StudentAnswerService:     service.NewStudentAnswerService(dbPool, repositories.StudentAnswerRepository),
		HomeworkResultService:    service.NewHomeworkResultService(dbPool, repositories.HomeworkResultRepository),
		TransactionService:       service.NewTransactionService(dbPool, repositories.TransactionRepository),
	}
}

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
	StudentAnswerService           service.StudentAnswerServiceInterface
	HomeworkResultService          service.HomeworkResultServiceInterface
	TransactionService             service.TransactionServiceInterface
	StudentCategoryStatsService    service.StudentCategoryStatsServiceInterface
	TeachersCoursesService         service.TeachersCoursesServiceInterface
	CourseLessonsService           service.CourseLessonsServiceInterface
	LessonsMaterialsService        service.LessonsMaterialsServiceInterface
	LessonHomeworksService         service.LessonHomeworksServiceInterface
	HomeworksTasksService          service.HomeworksTasksServiceInterface
	TransactionsCoursesService     service.TransactionsCoursesServiceInterface
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
		StudentAnswerService:           service.NewStudentAnswerService(dbPool, repositories.StudentAnswerRepository),
		HomeworkResultService:          service.NewHomeworkResultService(dbPool, repositories.HomeworkResultRepository),
		TransactionService:             service.NewTransactionService(dbPool, repositories.TransactionRepository),
		StudentCategoryStatsService:    service.NewStudentCategoryStatsService(dbPool, repositories.StudentCategoryStatsRepository),
		TeachersCoursesService:         service.NewTeachersCoursesService(dbPool, repositories.TeachersCoursesRepository),
		CourseLessonsService:           service.NewCourseLessonsService(dbPool, repositories.CourseLessonsRepository),
		LessonsMaterialsService:        service.NewLessonsMaterialsService(dbPool, repositories.LessonsMaterialsRepository),
		LessonHomeworksService:         service.NewLessonHomeworksService(dbPool, repositories.LessonHomeworksRepository),
		HomeworksTasksService:          service.NewHomeworksTasksService(dbPool, repositories.HomeworksTasksRepository),
		TransactionsCoursesService:     service.NewTransactionsCoursesService(dbPool, repositories.TransactionsCoursesRepository),
	}
}

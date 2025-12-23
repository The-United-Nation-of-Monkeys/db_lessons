package initializer

import "github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"

type RepositoryList struct {
	AuthRepository                repository.AuthRepositoryInterface
	StudentRepository             repository.StudentRepositoryInterface
	TeacherRepository             repository.TeacherRepositoryInterface
	CategoryRepository            repository.CategoryRepositoryInterface
	CurrencyRepository            repository.CurrencyRepositoryInterface
	LevelRepository               repository.LevelRepositoryInterface
	SubcategoryRepository         repository.SubcategoryRepositoryInterface
	MaterialRepository            repository.MaterialRepositoryInterface
	HomeworkRepository            repository.HomeworkRepositoryInterface
	LessonRepository              repository.LessonRepositoryInterface
	CourseRepository              repository.CourseRepositoryInterface
	TaskRepository                repository.TaskRepositoryInterface
	StatusHomeworkRepository      repository.StatusHomeworkRepositoryInterface
	StatusAnswerRepository        repository.StatusAnswerRepositoryInterface
	StatusTransactionRepository   repository.StatusTransactionRepositoryInterface
	StudentAnswerRepository       repository.StudentAnswerRepositoryInterface
	HomeworkResultRepository      repository.HomeworkResultRepositoryInterface
	TransactionRepository         repository.TransactionRepositoryInterface
	TeachersCoursesRepository     repository.TeachersCoursesRepositoryInterface
	CourseLessonsRepository       repository.CourseLessonsRepositoryInterface
	LessonsMaterialsRepository    repository.LessonsMaterialsRepositoryInterface
	LessonHomeworksRepository     repository.LessonHomeworksRepositoryInterface
	HomeworksTasksRepository           repository.HomeworksTasksRepositoryInterface
	TransactionsCoursesRepository      repository.TransactionsCoursesRepositoryInterface
	StudentCategoryStatsRepository     repository.StudentCategoryStatsRepositoryInterface
	SQLExecuteRepository               repository.SQLExecuteRepositoryInterface
}

func NewRepositoryList() *RepositoryList {
	return &RepositoryList{
		AuthRepository:                repository.NewAuthRepository(),
		StudentRepository:             repository.NewStudentRepository(),
		TeacherRepository:             repository.NewTeacherRepository(),
		CategoryRepository:            repository.NewCategoryRepository(),
		CurrencyRepository:            repository.NewCurrencyRepository(),
		LevelRepository:               repository.NewLevelRepository(),
		SubcategoryRepository:         repository.NewSubcategoryRepository(),
		MaterialRepository:            repository.NewMaterialRepository(),
		HomeworkRepository:            repository.NewHomeworkRepository(),
		LessonRepository:              repository.NewLessonRepository(),
		CourseRepository:              repository.NewCourseRepository(),
		TaskRepository:                repository.NewTaskRepository(),
		StatusHomeworkRepository:      repository.NewStatusHomeworkRepository(),
		StatusAnswerRepository:        repository.NewStatusAnswerRepository(),
		StatusTransactionRepository:   repository.NewStatusTransactionRepository(),
		StudentAnswerRepository:       repository.NewStudentAnswerRepository(),
		HomeworkResultRepository:      repository.NewHomeworkResultRepository(),
		TransactionRepository:         repository.NewTransactionRepository(),
		TeachersCoursesRepository:     repository.NewTeachersCoursesRepository(),
		CourseLessonsRepository:       repository.NewCourseLessonsRepository(),
		LessonsMaterialsRepository:    repository.NewLessonsMaterialsRepository(),
		LessonHomeworksRepository:     repository.NewLessonHomeworksRepository(),
		HomeworksTasksRepository:           repository.NewHomeworksTasksRepository(),
		TransactionsCoursesRepository:      repository.NewTransactionsCoursesRepository(),
		StudentCategoryStatsRepository:     repository.NewStudentCategoryStatsRepository(),
		SQLExecuteRepository:               repository.NewSQLExecuteRepository(),
	}
}

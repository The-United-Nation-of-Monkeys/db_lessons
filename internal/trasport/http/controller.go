package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Flussen/swagger-fiber-v3"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/docs"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/config"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/container/initializer"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/recover"
	redisstorage "github.com/gofiber/storage/redis/v3"
)

func NewController(server *fiber.App, cfg *config.Config, services *initializer.ServiceList, redisStg *redisstorage.Storage) {
	server.Use(recover.New())
	server.Use(logger.Middleware(&cfg.Logger))
	server.Use(cors.New())
	server.Use(helmet.New())
	server.Use(cache.New(cache.Config{
		Storage:      redisStg,
		Expiration:   10 * time.Second,
		CacheControl: true,
	}))

	api := server.Group(fmt.Sprintf("/api/v%d", cfg.Server.Version))
	api.Use("/swagger/*", swagger.HandlerDefault)
	docs.SwaggerInfo.Version = strconv.Itoa(cfg.Server.Version)
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/api/v%d", cfg.Server.Version)

	// Student routes
	studentHandler := NewStudentHandler(services.StudentService)
	api.Post("/students", studentHandler.Create)
	api.Get("/students", studentHandler.GetAll)
	api.Get("/students/:id", studentHandler.GetByID)
	api.Put("/students/:id", studentHandler.Update)
	api.Delete("/students/:id", studentHandler.Delete)

	// Teacher routes
	teacherHandler := NewTeacherHandler(services.TeacherService)
	api.Post("/teachers", teacherHandler.Create)
	api.Get("/teachers", teacherHandler.GetAll)
	api.Get("/teachers/:id", teacherHandler.GetByID)
	api.Put("/teachers/:id", teacherHandler.Update)
	api.Delete("/teachers/:id", teacherHandler.Delete)

	// Category routes
	categoryHandler := NewCategoryHandler(services.CategoryService)
	api.Post("/categories", categoryHandler.Create)
	api.Get("/categories", categoryHandler.GetAll)
	api.Get("/categories/:id", categoryHandler.GetByID)
	api.Put("/categories/:id", categoryHandler.Update)
	api.Delete("/categories/:id", categoryHandler.Delete)

	// Currency routes
	currencyHandler := NewCurrencyHandler(services.CurrencyService)
	api.Post("/currencies", currencyHandler.Create)
	api.Get("/currencies", currencyHandler.GetAll)
	api.Get("/currencies/:id", currencyHandler.GetByID)
	api.Put("/currencies/:id", currencyHandler.Update)
	api.Delete("/currencies/:id", currencyHandler.Delete)

	// Level routes
	levelHandler := NewLevelHandler(services.LevelService)
	api.Post("/levels", levelHandler.Create)
	api.Get("/levels", levelHandler.GetAll)
	api.Get("/levels/:id", levelHandler.GetByID)
	api.Put("/levels/:id", levelHandler.Update)
	api.Delete("/levels/:id", levelHandler.Delete)

	// Subcategory routes
	subcategoryHandler := NewSubcategoryHandler(services.SubcategoryService)
	api.Post("/subcategories", subcategoryHandler.Create)
	api.Get("/subcategories", subcategoryHandler.GetAll)
	api.Get("/subcategories/:id", subcategoryHandler.GetByID)
	api.Put("/subcategories/:id", subcategoryHandler.Update)
	api.Delete("/subcategories/:id", subcategoryHandler.Delete)

	// Material routes
	materialHandler := NewMaterialHandler(services.MaterialService)
	api.Post("/materials", materialHandler.Create)
	api.Get("/materials", materialHandler.GetAll)
	api.Get("/materials/:id", materialHandler.GetByID)
	api.Put("/materials/:id", materialHandler.Update)
	api.Delete("/materials/:id", materialHandler.Delete)

	// Homework routes
	homeworkHandler := NewHomeworkHandler(services.HomeworkService)
	api.Post("/homeworks", homeworkHandler.Create)
	api.Get("/homeworks", homeworkHandler.GetAll)
	api.Get("/homeworks/:id", homeworkHandler.GetByID)
	api.Put("/homeworks/:id", homeworkHandler.Update)
	api.Delete("/homeworks/:id", homeworkHandler.Delete)

	// Lesson routes
	lessonHandler := NewLessonHandler(services.LessonService)
	api.Post("/lessons", lessonHandler.Create)
	api.Get("/lessons", lessonHandler.GetAll)
	api.Get("/lessons/:id", lessonHandler.GetByID)
	api.Put("/lessons/:id", lessonHandler.Update)
	api.Delete("/lessons/:id", lessonHandler.Delete)

	// Course routes
	courseHandler := NewCourseHandler(services.CourseService)
	api.Post("/courses", courseHandler.Create)
	api.Get("/courses", courseHandler.GetAll)
	api.Get("/courses/:id", courseHandler.GetByID)
	api.Put("/courses/:id", courseHandler.Update)
	api.Delete("/courses/:id", courseHandler.Delete)

	// Task routes
	taskHandler := NewTaskHandler(services.TaskService)
	api.Post("/tasks", taskHandler.Create)
	api.Get("/tasks", taskHandler.GetAll)
	api.Get("/tasks/:id", taskHandler.GetByID)
	api.Put("/tasks/:id", taskHandler.Update)
	api.Delete("/tasks/:id", taskHandler.Delete)

	// StatusHomework routes
	statusHomeworkHandler := NewStatusHomeworkHandler(services.StatusHomeworkService)
	api.Post("/status-homeworks", statusHomeworkHandler.Create)
	api.Get("/status-homeworks", statusHomeworkHandler.GetAll)
	api.Get("/status-homeworks/:id", statusHomeworkHandler.GetByID)
	api.Put("/status-homeworks/:id", statusHomeworkHandler.Update)
	api.Delete("/status-homeworks/:id", statusHomeworkHandler.Delete)

	// StatusAnswer routes
	statusAnswerHandler := NewStatusAnswerHandler(services.StatusAnswerService)
	api.Post("/status-answers", statusAnswerHandler.Create)
	api.Get("/status-answers", statusAnswerHandler.GetAll)
	api.Get("/status-answers/:id", statusAnswerHandler.GetByID)
	api.Put("/status-answers/:id", statusAnswerHandler.Update)
	api.Delete("/status-answers/:id", statusAnswerHandler.Delete)

	// StatusTransaction routes
	statusTransactionHandler := NewStatusTransactionHandler(services.StatusTransactionService)
	api.Post("/status-transactions", statusTransactionHandler.Create)
	api.Get("/status-transactions", statusTransactionHandler.GetAll)
	api.Get("/status-transactions/:id", statusTransactionHandler.GetByID)
	api.Put("/status-transactions/:id", statusTransactionHandler.Update)
	api.Delete("/status-transactions/:id", statusTransactionHandler.Delete)

	// StudentAnswer routes
	studentAnswerHandler := NewStudentAnswerHandler(services.StudentAnswerService)
	api.Post("/student-answers", studentAnswerHandler.Create)
	api.Get("/student-answers", studentAnswerHandler.GetAll)
	api.Get("/student-answers/:id", studentAnswerHandler.GetByID)
	api.Put("/student-answers/:id", studentAnswerHandler.Update)
	api.Delete("/student-answers/:id", studentAnswerHandler.Delete)

	// HomeworkResult routes
	homeworkResultHandler := NewHomeworkResultHandler(services.HomeworkResultService)
	api.Post("/homework-results", homeworkResultHandler.Create)
	api.Get("/homework-results", homeworkResultHandler.GetAll)
	api.Get("/homework-results/:id", homeworkResultHandler.GetByID)
	api.Put("/homework-results/:id", homeworkResultHandler.Update)
	api.Delete("/homework-results/:id", homeworkResultHandler.Delete)

	// Transaction routes
	transactionHandler := NewTransactionHandler(services.TransactionService)
	api.Post("/transactions", transactionHandler.Create)
	api.Get("/transactions", transactionHandler.GetAll)
	// Специфичные роуты должны быть определены ПЕРЕД параметризованными
	api.Get("/transactions/report", transactionHandler.GetReportByParams)
	api.Post("/transactions/bulk-update-status", transactionHandler.BulkUpdateTransactionStatus)
	api.Get("/transactions/:id", transactionHandler.GetByID)
	api.Put("/transactions/:id", transactionHandler.Update)
	api.Delete("/transactions/:id", transactionHandler.Delete)

	// Student Category Stats routes (reports)
	studentCategoryStatsHandler := NewStudentCategoryStatsHandler(services.StudentCategoryStatsService)
	api.Get("/reports/student-category-stats", studentCategoryStatsHandler.GetAll)
	api.Get("/reports/student-category-stats/students/:id", studentCategoryStatsHandler.GetByStudentID)
	api.Get("/reports/student-category-stats/categories/:id", studentCategoryStatsHandler.GetByCategoryID)

	// Teachers-Courses relations routes
	teachersCoursesHandler := NewTeachersCoursesHandler(services.TeachersCoursesService)
	api.Post("/teachers-courses", teachersCoursesHandler.Create)
	api.Get("/teachers-courses", teachersCoursesHandler.GetAll)
	api.Get("/teachers-courses/:course_id/:teacher_id", teachersCoursesHandler.GetByID)
	api.Delete("/teachers-courses/:course_id/:teacher_id", teachersCoursesHandler.Delete)

	// Course-Lessons relations routes
	courseLessonsHandler := NewCourseLessonsHandler(services.CourseLessonsService)
	api.Post("/course-lessons", courseLessonsHandler.Create)
	api.Get("/course-lessons", courseLessonsHandler.GetAll)
	api.Get("/course-lessons/:id", courseLessonsHandler.GetByID)
	api.Delete("/course-lessons/:id", courseLessonsHandler.Delete)

	// Lessons-Materials relations routes
	lessonsMaterialsHandler := NewLessonsMaterialsHandler(services.LessonsMaterialsService)
	api.Post("/lessons-materials", lessonsMaterialsHandler.Create)
	api.Get("/lessons-materials", lessonsMaterialsHandler.GetAll)
	api.Get("/lessons-materials/:lesson_id/:material_id", lessonsMaterialsHandler.GetByID)
	api.Delete("/lessons-materials/:lesson_id/:material_id", lessonsMaterialsHandler.Delete)

	// Lesson-Homeworks relations routes
	lessonHomeworksHandler := NewLessonHomeworksHandler(services.LessonHomeworksService)
	api.Post("/lesson-homeworks", lessonHomeworksHandler.Create)
	api.Get("/lesson-homeworks", lessonHomeworksHandler.GetAll)
	api.Get("/lesson-homeworks/:lesson_id/:homework_id", lessonHomeworksHandler.GetByID)
	api.Delete("/lesson-homeworks/:lesson_id/:homework_id", lessonHomeworksHandler.Delete)

	// Homeworks-Tasks relations routes
	homeworksTasksHandler := NewHomeworksTasksHandler(services.HomeworksTasksService)
	api.Post("/homeworks-tasks", homeworksTasksHandler.Create)
	api.Get("/homeworks-tasks", homeworksTasksHandler.GetAll)
	api.Get("/homeworks-tasks/:task_id/:homework_id", homeworksTasksHandler.GetByID)
	api.Delete("/homeworks-tasks/:task_id/:homework_id", homeworksTasksHandler.Delete)

	// Transactions-Courses relations routes
	transactionsCoursesHandler := NewTransactionsCoursesHandler(services.TransactionsCoursesService)
	api.Post("/transactions-courses", transactionsCoursesHandler.Create)
	api.Get("/transactions-courses", transactionsCoursesHandler.GetAll)
	api.Get("/transactions-courses/:transaction_id/:course_id", transactionsCoursesHandler.GetByID)
	api.Delete("/transactions-courses/:transaction_id/:course_id", transactionsCoursesHandler.Delete)

	api.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.Status(200).SendString("pong")
	})
}

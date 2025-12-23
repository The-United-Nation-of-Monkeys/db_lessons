package initializer

import (
	"time"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/config"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/jwt"
)

type ServiceList struct {
	JWTService                  *jwt.ServiceJWT
	AuthService                 service.AuthServiceInterface
	StudentService              service.StudentServiceInterface
	TeacherService              service.TeacherServiceInterface
	CategoryService             service.CategoryServiceInterface
	CurrencyService             service.CurrencyServiceInterface
	LevelService                service.LevelServiceInterface
	SubcategoryService          service.SubcategoryServiceInterface
	MaterialService             service.MaterialServiceInterface
	HomeworkService             service.HomeworkServiceInterface
	LessonService               service.LessonServiceInterface
	CourseService               service.CourseServiceInterface
	TaskService                 service.TaskServiceInterface
	StatusHomeworkService       service.StatusHomeworkServiceInterface
	StatusAnswerService         service.StatusAnswerServiceInterface
	StatusTransactionService    service.StatusTransactionServiceInterface
	StudentAnswerService        service.StudentAnswerServiceInterface
	HomeworkResultService       service.HomeworkResultServiceInterface
	TransactionService          service.TransactionServiceInterface
	StudentCategoryStatsService service.StudentCategoryStatsServiceInterface
	TeachersCoursesService      service.TeachersCoursesServiceInterface
	CourseLessonsService        service.CourseLessonsServiceInterface
	LessonsMaterialsService     service.LessonsMaterialsServiceInterface
	LessonHomeworksService      service.LessonHomeworksServiceInterface
	HomeworksTasksService       service.HomeworksTasksServiceInterface
	TransactionsCoursesService  service.TransactionsCoursesServiceInterface
	SQLExecuteService           service.SQLExecuteServiceInterface
}

func NewServiceList(repositories *RepositoryList, cfg *config.Config) *ServiceList {
	// Load JWT keys
	privateKeyPath := cfg.JWT.GetPrivateKeyPath()
	publicKeyPath := cfg.JWT.GetPublicKeyPath()

	privateKey, err := jwt.LoadPrivateKey(privateKeyPath)
	if err != nil {
		panic("failed to load private key: " + err.Error())
	}

	publicKey, err := jwt.LoadPublicKey(publicKeyPath)
	if err != nil {
		panic("failed to load public key: " + err.Error())
	}

	// Parse time durations
	refreshTimeExp, err := time.ParseDuration(cfg.JWT.RefreshTimeExp)
	if err != nil {
		panic("failed to parse refresh time exp: " + err.Error())
	}

	accessTimeExp, err := time.ParseDuration(cfg.JWT.AccessTimeExp)
	if err != nil {
		panic("failed to parse access time exp: " + err.Error())
	}

	// Create JWT service
	jwtService := jwt.NewServiceJWT(privateKey, publicKey, refreshTimeExp, accessTimeExp)

	// Create BaseService with role mappings
	roleMap := map[string]service.RoleConfig{
		"admin":   {User: "app_admin_user"},
		"teacher": {User: "app_teacher_user"},
		"student": {User: "app_student_user"},
	}
	baseService := service.NewBaseService(roleMap, cfg.DataBase, "app_base_user")

	// Create auth service - использует базового пользователя через BaseService
	authService := service.NewAuthService(repositories.AuthRepository, jwtService, baseService)

	return &ServiceList{
		JWTService:                  jwtService,
		AuthService:                 authService,
		StudentService:              service.NewStudentService(baseService, repositories.StudentRepository),
		TeacherService:              service.NewTeacherService(baseService, repositories.TeacherRepository),
		CategoryService:             service.NewCategoryService(baseService, repositories.CategoryRepository),
		CurrencyService:             service.NewCurrencyService(baseService, repositories.CurrencyRepository),
		LevelService:                service.NewLevelService(baseService, repositories.LevelRepository),
		SubcategoryService:          service.NewSubcategoryService(baseService, repositories.SubcategoryRepository),
		MaterialService:             service.NewMaterialService(baseService, repositories.MaterialRepository),
		HomeworkService:             service.NewHomeworkService(baseService, repositories.HomeworkRepository),
		LessonService:               service.NewLessonService(baseService, repositories.LessonRepository),
		CourseService:               service.NewCourseService(baseService, repositories.CourseRepository),
		TaskService:                 service.NewTaskService(baseService, repositories.TaskRepository),
		StatusHomeworkService:       service.NewStatusHomeworkService(baseService, repositories.StatusHomeworkRepository),
		StatusAnswerService:         service.NewStatusAnswerService(baseService, repositories.StatusAnswerRepository),
		StatusTransactionService:    service.NewStatusTransactionService(baseService, repositories.StatusTransactionRepository),
		StudentAnswerService:        service.NewStudentAnswerService(baseService, repositories.StudentAnswerRepository),
		HomeworkResultService:       service.NewHomeworkResultService(baseService, repositories.HomeworkResultRepository),
		TransactionService:          service.NewTransactionService(baseService, repositories.TransactionRepository),
		StudentCategoryStatsService: service.NewStudentCategoryStatsService(baseService, repositories.StudentCategoryStatsRepository),
		TeachersCoursesService:      service.NewTeachersCoursesService(baseService, repositories.TeachersCoursesRepository),
		CourseLessonsService:        service.NewCourseLessonsService(baseService, repositories.CourseLessonsRepository),
		LessonsMaterialsService:     service.NewLessonsMaterialsService(baseService, repositories.LessonsMaterialsRepository),
		LessonHomeworksService:      service.NewLessonHomeworksService(baseService, repositories.LessonHomeworksRepository),
		HomeworksTasksService:       service.NewHomeworksTasksService(baseService, repositories.HomeworksTasksRepository),
		TransactionsCoursesService:  service.NewTransactionsCoursesService(baseService, repositories.TransactionsCoursesRepository),
		SQLExecuteService:           service.NewSQLExecuteService(baseService, repositories.SQLExecuteRepository),
	}
}

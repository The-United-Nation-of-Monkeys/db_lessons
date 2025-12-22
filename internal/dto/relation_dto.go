package dto

// TeachersCoursesDTO represents the many-to-many relationship between teachers and courses
type CreateTeachersCoursesDTO struct {
	CourseID  int `json:"course_id" validate:"required"`
	TeacherID int `json:"teacher_id" validate:"required"`
}

type TeachersCoursesDTO struct {
	CourseID  int `json:"course_id"`
	TeacherID int `json:"teacher_id"`
}

// CourseLessonsDTO represents the relationship between courses and lessons
type CreateCourseLessonsDTO struct {
	LessonID int `json:"lesson_id" validate:"required"`
	CourseID int `json:"course_id" validate:"required"`
}

type CourseLessonsDTO struct {
	LessonID int `json:"lesson_id"`
	CourseID int `json:"course_id"`
}

// LessonsMaterialsDTO represents the many-to-many relationship between lessons and materials
type CreateLessonsMaterialsDTO struct {
	LessonID   int `json:"lesson_id" validate:"required"`
	MaterialID int `json:"material_id" validate:"required"`
}

type LessonsMaterialsDTO struct {
	LessonID   int `json:"lesson_id"`
	MaterialID int `json:"material_id"`
}

// LessonHomeworksDTO represents the many-to-many relationship between lessons and homeworks
type CreateLessonHomeworksDTO struct {
	LessonID   int `json:"lesson_id" validate:"required"`
	HomeworkID int `json:"homework_id" validate:"required"`
}

type LessonHomeworksDTO struct {
	LessonID   int `json:"lesson_id"`
	HomeworkID int `json:"homework_id"`
}

// HomeworksTasksDTO represents the many-to-many relationship between homeworks and tasks
type CreateHomeworksTasksDTO struct {
	TaskID     int `json:"task_id" validate:"required"`
	HomeworkID int `json:"homework_id" validate:"required"`
}

type HomeworksTasksDTO struct {
	TaskID     int `json:"task_id"`
	HomeworkID int `json:"homework_id"`
}

// TransactionsCoursesDTO represents the many-to-many relationship between transactions and courses
type CreateTransactionsCoursesDTO struct {
	TransactionID int `json:"transaction_id" validate:"required"`
	CourseID      int `json:"course_id" validate:"required"`
}

type TransactionsCoursesDTO struct {
	TransactionID int `json:"transaction_id"`
	CourseID      int `json:"course_id"`
}



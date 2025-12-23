package dto

type CreateTeachersCoursesDTO struct {
	CourseID  int `json:"course_id" validate:"required"`
	TeacherID int `json:"teacher_id" validate:"required"`
}

type TeachersCoursesDTO struct {
	CourseID  int `json:"course_id"`
	TeacherID int `json:"teacher_id"`
}

type CreateCourseLessonsDTO struct {
	LessonID int `json:"lesson_id" validate:"required"`
	CourseID int `json:"course_id" validate:"required"`
}

type CourseLessonsDTO struct {
	LessonID int `json:"lesson_id"`
	CourseID int `json:"course_id"`
}

type CreateLessonsMaterialsDTO struct {
	LessonID   int `json:"lesson_id" validate:"required"`
	MaterialID int `json:"material_id" validate:"required"`
}

type LessonsMaterialsDTO struct {
	LessonID   int `json:"lesson_id"`
	MaterialID int `json:"material_id"`
}

type CreateLessonHomeworksDTO struct {
	LessonID   int `json:"lesson_id" validate:"required"`
	HomeworkID int `json:"homework_id" validate:"required"`
}

type LessonHomeworksDTO struct {
	LessonID   int `json:"lesson_id"`
	HomeworkID int `json:"homework_id"`
}

type CreateHomeworksTasksDTO struct {
	TaskID     int `json:"task_id" validate:"required"`
	HomeworkID int `json:"homework_id" validate:"required"`
}

type HomeworksTasksDTO struct {
	TaskID     int `json:"task_id"`
	HomeworkID int `json:"homework_id"`
}

type CreateTransactionsCoursesDTO struct {
	TransactionID int `json:"transaction_id" validate:"required"`
	CourseID      int `json:"course_id" validate:"required"`
}

type TransactionsCoursesDTO struct {
	TransactionID int `json:"transaction_id"`
	CourseID      int `json:"course_id"`
}



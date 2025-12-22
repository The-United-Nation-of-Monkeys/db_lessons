package dto

type CreateStudentAnswerDTO struct {
	StudentID      int    `json:"student_id" validate:"required"`
	TaskID         int    `json:"task_id" validate:"required"`
	Answer         string `json:"answer"`
	StatusAnswerID *int   `json:"status_answer_id"`
}

type UpdateStudentAnswerDTO struct {
	StudentID      *int    `json:"student_id"`
	TaskID         *int    `json:"task_id"`
	Answer         *string `json:"answer"`
	StatusAnswerID *int    `json:"status_answer_id"`
}

type StudentAnswerDTO struct {
	StudentAnswerID int    `json:"student_answer_id"`
	StudentID        int    `json:"student_id"`
	TaskID           int    `json:"task_id"`
	Answer           string `json:"answer"`
	StatusAnswerID   *int   `json:"status_answer_id"`
}



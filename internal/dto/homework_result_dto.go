package dto

type CreateHomeworkResultDTO struct {
	StudentID        int    `json:"student_id" validate:"required"`
	TaskID           int    `json:"task_id" validate:"required"`
	Answer           string `json:"answer"`
	StatusHomeworkID *int   `json:"status_homework_id"`
	Points           int    `json:"points" validate:"gte=0"`
}

type UpdateHomeworkResultDTO struct {
	StudentID        *int    `json:"student_id"`
	TaskID           *int    `json:"task_id"`
	Answer           *string `json:"answer"`
	StatusHomeworkID *int    `json:"status_homework_id"`
	Points           *int    `json:"points" validate:"omitempty,gte=0"`
}

type HomeworkResultDTO struct {
	StudentAnswerID  int    `json:"student_answer_id"`
	StudentID        int    `json:"student_id"`
	TaskID           int    `json:"task_id"`
	Answer           string `json:"answer"`
	StatusHomeworkID *int   `json:"status_homework_id"`
	Points           int    `json:"points"`
}



package dto

type CreateHomeworkResultDTO struct {
	StudentAnswerID  int  `json:"student_answer_id" validate:"required"`
	StudentID        int  `json:"student_id" validate:"required"`
	TaskID           int  `json:"task_id" validate:"required"`
	StatusHomeworkID *int `json:"status_homework_id"`
	Points           int  `json:"points" validate:"gte=0"`
}

type UpdateHomeworkResultDTO struct {
	StudentAnswerID  *int `json:"student_answer_id"`
	StudentID        *int `json:"student_id"`
	TaskID           *int `json:"task_id"`
	StatusHomeworkID *int `json:"status_homework_id"`
	Points           *int `json:"points" validate:"omitempty,gte=0"`
}

type HomeworkResultDTO struct {
	HomeworkResultID int    `json:"homework_result_id"`
	StudentAnswerID  int    `json:"student_answer_id"`
	StudentID        int    `json:"student_id"`
	TaskID           int    `json:"task_id"`
	StatusHomeworkID *int   `json:"status_homework_id"`
	Points           int    `json:"points"`
}



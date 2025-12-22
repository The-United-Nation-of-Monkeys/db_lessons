package dto

import "time"

type CreateHomeworkDTO struct {
	Name         string     `json:"name" validate:"required"`
	Description  string     `json:"description"`
	DeadlineTime *time.Time `json:"deadline_time"`
}

type UpdateHomeworkDTO struct {
	Name         *string    `json:"name"`
	Description  *string    `json:"description"`
	DeadlineTime *time.Time `json:"deadline_time"`
}

type HomeworkDTO struct {
	HomeworkID   int        `json:"homework_id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	DeadlineTime *time.Time `json:"deadline_time"`
}



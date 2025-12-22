package dto

type CreateStatusAnswerDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateStatusAnswerDTO struct {
	Name *string `json:"name"`
}

type StatusAnswerDTO struct {
	StatusAnswerID int    `json:"status_answer_id"`
	Name           string `json:"name"`
}



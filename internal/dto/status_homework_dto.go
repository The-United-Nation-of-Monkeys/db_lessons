package dto

type CreateStatusHomeworkDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateStatusHomeworkDTO struct {
	Name *string `json:"name"`
}

type StatusHomeworkDTO struct {
	StatusHomeworkID int    `json:"status_homework_id"`
	Name             string `json:"name"`
}



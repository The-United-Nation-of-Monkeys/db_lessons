package dto

type CreateLevelDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateLevelDTO struct {
	Name *string `json:"name"`
}

type LevelDTO struct {
	LevelID int    `json:"level_id"`
	Name    string `json:"name"`
}



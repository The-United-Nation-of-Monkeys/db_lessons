package dto

type CreateTaskDTO struct {
	Type          string `json:"type" validate:"required"`
	Description   string `json:"description"`
	RightAnswer   string `json:"right_answer"`
	Points        int    `json:"points" validate:"gte=0"`
	LevelID       *int   `json:"level_id"`
	CategoryID    *int   `json:"category_id"`
	SubcategoryID *int   `json:"subcategory_id"`
}

type UpdateTaskDTO struct {
	Type          *string `json:"type"`
	Description   *string `json:"description"`
	RightAnswer   *string `json:"right_answer"`
	Points        *int    `json:"points" validate:"omitempty,gte=0"`
	LevelID       *int    `json:"level_id"`
	CategoryID    *int    `json:"category_id"`
	SubcategoryID *int    `json:"subcategory_id"`
}

type TaskDTO struct {
	TaskID        int    `json:"task_id"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	RightAnswer   string `json:"right_answer"`
	Points        int    `json:"points"`
	LevelID       *int   `json:"level_id"`
	CategoryID    *int   `json:"category_id"`
	SubcategoryID *int   `json:"subcategory_id"`
}



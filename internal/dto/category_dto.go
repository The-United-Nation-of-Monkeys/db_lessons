package dto

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name"`
}

type CategoryDTO struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}



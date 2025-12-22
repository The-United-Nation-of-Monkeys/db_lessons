package dto

type CreateSubcategoryDTO struct {
	Name       string `json:"name" validate:"required"`
	CategoryID *int   `json:"category_id"`
}

type UpdateSubcategoryDTO struct {
	Name       *string `json:"name"`
	CategoryID *int    `json:"category_id"`
}

type SubcategoryDTO struct {
	SubcategoryID int    `json:"subcategory_id"`
	Name          string `json:"name"`
	CategoryID    *int   `json:"category_id"`
}



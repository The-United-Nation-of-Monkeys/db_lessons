package dto

type CreateMaterialDTO struct {
	Source    string `json:"source" validate:"required"`
	Extension string `json:"extension" validate:"required,max=7"`
	Size      int    `json:"size" validate:"gte=0"`
}

type UpdateMaterialDTO struct {
	Source    *string `json:"source"`
	Extension *string `json:"extension" validate:"omitempty,max=7"`
	Size      *int    `json:"size" validate:"omitempty,gte=0"`
}

type MaterialDTO struct {
	MaterialID int    `json:"material_id"`
	Source     string `json:"source"`
	Extension  string `json:"extension"`
	Size       int    `json:"size"`
}



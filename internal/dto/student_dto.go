package dto

type CreateStudentDTO struct {
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	BonusAmount int    `json:"bonus_amount" validate:"gte=0"`
}

type UpdateStudentDTO struct {
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Email       *string `json:"email" validate:"omitempty,email"`
	Password    *string `json:"password"`
	BonusAmount *int    `json:"bonus_amount" validate:"omitempty,gte=0"`
}

type StudentDTO struct {
	StudentID   int    `json:"student_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BonusAmount int    `json:"bonus_amount"`
}



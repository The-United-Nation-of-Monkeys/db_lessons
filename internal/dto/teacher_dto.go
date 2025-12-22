package dto

type CreateTeacherDTO struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateTeacherDTO struct {
	Name     *string `json:"name"`
	Surname  *string `json:"surname"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password"`
}

type TeacherDTO struct {
	TeacherID int    `json:"teacher_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}



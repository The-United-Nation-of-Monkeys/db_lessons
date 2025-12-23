package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin student teacher"`
}

type AuthResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int    `json:"user_id"`
	Role         string `json:"role"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token"`
}

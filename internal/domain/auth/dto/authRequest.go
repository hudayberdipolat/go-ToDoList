package dto

type RegisterRequest struct {
	Username        string `json:"username"  validate:"required,min=3"`
	FullName        string `json:"full_name"  validate:"required,min=3"`
	Password        string `json:"password"  validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=3"`
}

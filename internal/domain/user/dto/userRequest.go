package dto

type UpdateUserData struct {
	Username string `json:"username" validate:"required,min=3"`
	FullName string `json:"full_name" validate:"required,min=3"`
}

type UpdateUserPassword struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

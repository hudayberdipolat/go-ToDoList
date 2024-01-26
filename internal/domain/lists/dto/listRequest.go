package dto

type CreateListRequest struct {
	ListName string `json:"list_name" validate:"required"`
}

type UpdateListRequest struct {
	ListName   string `json:"list_name" validate:"required"`
	ListStatus string `json:"list_status" validate:"required"`
}

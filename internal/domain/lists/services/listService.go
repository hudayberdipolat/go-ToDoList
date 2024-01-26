package services

import "github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/dto"

type ListService interface {
	GetListsOfUser(userID int) ([]dto.ListResponse, error)
	GetListOfUser(userID, listID int) (*dto.ListResponse, error)
	CreateList(userID int, createListRequest dto.CreateListRequest) error
	UpdateList(userID, listID int, updateListRequest dto.UpdateListRequest) error
	DeleteList(userID, listID int) error
}

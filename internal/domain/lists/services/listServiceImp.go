package services

import (
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/repositories"
	"github.com/hudayberdipolat/go-ToDoList/internal/models"
	"time"
)

type listServiceImp struct {
	listRepo repositories.ListRepository
}

func NewListService(repo repositories.ListRepository) ListService {
	return listServiceImp{
		listRepo: repo,
	}
}

func (l listServiceImp) GetListsOfUser(userID int) ([]dto.ListResponse, error) {
	lists, err := l.listRepo.FindAll(userID)
	if err != nil {
		return nil, err
	}
	var listResponses []dto.ListResponse
	for _, list := range lists {
		listResponse := dto.NewListResponse(list)
		listResponses = append(listResponses, listResponse)
	}
	return listResponses, nil
}

func (l listServiceImp) GetListOfUser(userID, listID int) (*dto.ListResponse, error) {
	list, err := l.listRepo.FindOne(userID, listID)
	if err != nil {
		return nil, err
	}
	listResponse := dto.NewListResponse(*list)
	return &listResponse, nil
}

func (l listServiceImp) CreateList(userID int, createListRequest dto.CreateListRequest) error {
	createList := models.List{
		ListName:   createListRequest.ListName,
		ListStatus: "1",
		UserID:     uint(userID),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := l.listRepo.Create(createList); err != nil {
		return err
	}
	return nil
}

func (l listServiceImp) UpdateList(userID, listID int, updateListRequest dto.UpdateListRequest) error {
	updateList := models.List{
		ListName:   updateListRequest.ListName,
		ListStatus: updateListRequest.ListStatus,
		UpdatedAt:  time.Now(),
	}
	if err := l.listRepo.Update(userID, listID, updateList); err != nil {
		return err
	}
	return nil
}

func (l listServiceImp) DeleteList(userID, listID int) error {

	if err := l.listRepo.Delete(userID, listID); err != nil {
		return err
	}
	return nil
}

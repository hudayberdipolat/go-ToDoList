package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-ToDoList/internal/domain/user/services"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/response"
	"github.com/hudayberdipolat/go-ToDoList/internal/utils/validate"
	"net/http"
)

type userHandlerImp struct {
	userService services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}

}

func (u userHandlerImp) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	user, err := u.userService.GetUserData(userID)
	if err != nil {
		errResponse := response.Error(http.StatusNotFound, "user not found", err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get user data", user)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserData(ctx *fiber.Ctx) error {
	var updateDataRequest dto.UpdateUserData
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&updateDataRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateDataRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	// update user Data
	if err := u.userService.UpdateUserData(userID, updateDataRequest); err != nil {
		errResponse := response.Error(http.StatusInternalServerError, "user data can't updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "update user profile successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserPassword(ctx *fiber.Ctx) error {
	var updatePasswordRequest dto.UpdateUserPassword
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&updatePasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updatePasswordRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.userService.UpdateUserPassword(userID, updatePasswordRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user data can't updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "changed user password successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

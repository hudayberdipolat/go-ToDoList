package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	GetUser(ctx *fiber.Ctx) error
	UpdateUserData(ctx *fiber.Ctx) error
	UpdateUserPassword(ctx *fiber.Ctx) error
}

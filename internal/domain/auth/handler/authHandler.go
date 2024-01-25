package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

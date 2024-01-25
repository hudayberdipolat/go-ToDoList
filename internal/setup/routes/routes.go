package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/constructor"
)

func Routes(router *fiber.App) {
	apiRoute := router.Group("/api")

	// auth user route
	authRoute := apiRoute.Group("/auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)
}

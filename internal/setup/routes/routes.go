package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/constructor"
	userConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/user/constructor"
	"github.com/hudayberdipolat/go-ToDoList/internal/middleware"
)

func Routes(router *fiber.App) {
	apiRoute := router.Group("/api")

	// auth user route
	authRoute := apiRoute.Group("/auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)

	// user profile data
	userRoute := apiRoute.Group("/user")
	userRoute.Use(middleware.AuthMiddleware)
	userRoute.Get("", userConstructor.UserHandler.GetUser)
	userRoute.Put("/update-data", userConstructor.UserHandler.UpdateUserData)
	userRoute.Put("/update-password", userConstructor.UserHandler.UpdateUserPassword)
}

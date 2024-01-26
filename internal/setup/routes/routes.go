package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/auth/constructor"
	listConstructor "github.com/hudayberdipolat/go-ToDoList/internal/domain/lists/constructor"
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

	// list routes
	listRoute := userRoute.Group("/lists")
	listRoute.Get("", listConstructor.ListHandler.GetAll)
	listRoute.Get("/:listID", listConstructor.ListHandler.GetOne)
	listRoute.Post("/create", listConstructor.ListHandler.Create)
	listRoute.Put("/:listID/update", listConstructor.ListHandler.Update)
	listRoute.Delete("/:listID/delete", listConstructor.ListHandler.Delete)
}

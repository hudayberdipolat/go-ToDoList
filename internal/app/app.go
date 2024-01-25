package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-ToDoList/internal/setup/routes"
	"time"
)

func NewApp(dependencies *Dependencies) (httpServer *fiber.App) {
	httpServer = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"status":  code,
				"message": "Näsazlyk ýüze çykdy, Sonrak synanysyn!!!",
			})
		},
		AppName:      dependencies.Config.HttpServer.AppName,
		BodyLimit:    30 * 1024 * 1024,
		ServerHeader: dependencies.Config.HttpServer.AppHeader,
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	})
	routes.Routes(httpServer)

	return httpServer
}

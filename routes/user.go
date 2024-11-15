package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wignn/go-with-mongoDb/service"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", service.GetUsers)
	app.Post("/users", service.CreateUser)
	app.Put("/users/:id", service.UpdateUser)
	app.Delete("/users/:id", service.DeleteUser)
}

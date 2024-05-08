package employee

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Post("/employees/create", createHandler)
	app.Put("/employees/update", updateHandler)
	app.Get("/employees/getByID", getByIDHandler)
	app.Delete("/employees/deleteByID", deleteByIDHandler)
	app.Get("/employees/list", listHandler)
}

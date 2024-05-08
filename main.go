package main

import (
	"employee-go/employee"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	employee.Init()
	employee.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("something went wrong", err)
	}
}

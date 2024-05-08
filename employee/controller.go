package employee

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func createHandler(c *fiber.Ctx) error {
	var employee struct {
		Name     string  `json:"name"`
		Position string  `json:"position"`
		Salary   float64 `json:"salary"`
	}
	if err := c.BodyParser(&employee); err != nil {
		return err
	}
	emp, err := Store.create(employee.Name, employee.Position, employee.Salary)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(emp)
}

func updateHandler(c *fiber.Ctx) error {
	employee := Employee{}
	if err := c.BodyParser(&employee); err != nil {
		return err
	}
	emp, err := Store.update(employee.ID, employee.Name, employee.Position, employee.Salary)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(emp)
}

func getByIDHandler(c *fiber.Ctx) error {
	var request struct {
		ID int `json:"id"`
	}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	emp, err := Store.readByID(request.ID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(emp)
}

func deleteByIDHandler(c *fiber.Ctx) error {
	var request struct {
		ID int `json:"id"`
	}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	err := Store.deleteByID(request.ID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "successfully deleted"})
}

func listHandler(c *fiber.Ctx) error {
	var request struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	employees, err := Store.list(request.Page, request.Limit)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(employees)
}

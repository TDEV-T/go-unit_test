package main

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Add(x, y int) int {
	return x + y
}

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}

var validate = validator.New()

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,fullname"`
	Age      int    `json:"age" validate:"required,numeric,min=1"`
}

func setup() *fiber.App {
	app := fiber.New()

	validate.RegisterValidation("fullname", validateFullname)

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON "})
		}

		if err := validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	return app
}

func main() {
	app := setup()
	app.Listen(":8000")

}

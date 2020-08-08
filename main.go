package main

import (
	"github.com/allanfvc/uni7sum/controller"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type customError struct {
	Title   string `json: "title"`
	Message string `json: "message`
}

func main() {
	app := fiber.New()
	app.Settings.ErrorHandler = errorHandler
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	controller.RegisterRoutes(app)
	app.Listen(8080)
}

func errorHandler(ctx *fiber.Ctx, err error) {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	custom := customError{
		Title:   "An unexpected error occurred",
		Message: err.Error(),
	}

	ctx.Status(code).JSON(custom)
}

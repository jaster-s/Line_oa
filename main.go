package main

import (
	"fli-gateway-api/lib"
	"fli-gateway-api/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func main() {
	_logger := new(lib.Logger)
	PORT := os.Getenv("APP_PORT")
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Register Router
	router.New(app)

	err := app.Listen(fmt.Sprintf(":%s", PORT))

	if err != nil {
		_logger.Error(fmt.Sprintf("Cannot start server: %v", err), true)
	}
}

package router

import (
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Freshdesk Line Integration Gateway API")
	})

	// Line OA Router
	lineOA := app.Group("/line_oa")
	LineOA(lineOA)
}
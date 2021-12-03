package controller

import (
	"fli-gateway-api/lib"
	"fli-gateway-api/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type LineOAController struct {}

func (LineOAController) ListLineOA(c *fiber.Ctx) error {
	_logger := new(lib.Logger)
	lineOAModel := new(models.LineOAModel)

	lineOA, err := lineOAModel.List()

	if err != nil {
		_logger.Error(fmt.Sprintf("%v", err), false)

		return c.Status(500).JSON(map[string]string {
			"code": "unknown-error",
			"message": "Something went wrong. Please try again.",
		})
	}

	return c.JSON(lineOA)
}

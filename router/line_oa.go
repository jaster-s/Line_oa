package router

import (
	"fli-gateway-api/controller"
	"github.com/gofiber/fiber/v2"
)

func LineOA(router fiber.Router) {
	lineOACtrl := new(controller.LineOAController)

	router.Get("/list", lineOACtrl.ListLineOA)
}
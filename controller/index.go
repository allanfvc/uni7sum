package controller

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
)

//RegisterRoutes register all app routes
func RegisterRoutes(app fiber.Router) {
	app.Get("/sum", sum)
}

func sum(ctx *fiber.Ctx) {
	firstParam, errFirstParam := strconv.ParseFloat(ctx.Query("a"), 32)
	secondParam, errSecondParam := strconv.ParseFloat(ctx.Query("b"), 32)

	if errFirstParam == nil && errSecondParam == nil {
		ctx.Status(200).Send(firstParam + secondParam)
		return
	}
	ctx.Status(400).Send(fmt.Sprintf("Error on trying to sum %s and %s", ctx.Query("a"), ctx.Query("b")))
}

package router

import (
	"app/server/api"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				return ctx.JSON(fiber.Map{
					"code":    201,
					"message": err.Error(),
				})
			}
			return nil
		},
	})

	app.Use(recover2.New())

	var bc api.BookController

	app.Get("/search", bc.Search)

	app.Get("/info", bc.Info)

	app.Get("/body", bc.Body)

	return app
}

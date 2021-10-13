package router

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	rec "github.com/gofiber/fiber/v2/middleware/recover"
)

func New(config ...fiber.Config) *fiber.App {

	app := fiber.New(config...)

	app.Use(rec.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"code":200,
			"message":"ok",
		})
	})

	app.Get("/err", func(ctx *fiber.Ctx) error {
		panic(errors.New("出错啦"))
	})

	return app
}

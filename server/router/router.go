package router

import (
	"app/server/common/result"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func New(config ...fiber.Config) *fiber.App {

	app := fiber.New(config...)

	//app.Use(rec.New())

	var r result.Result

	app.Get("/", func(ctx *fiber.Ctx) error {
		return r.Ctx(ctx).Success()
	})

	app.Get("/err", func(ctx *fiber.Ctx) error {
		panic(errors.New("出错啦"))
	})

	return app
}

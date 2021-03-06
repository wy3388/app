package router

import (
	"app/server/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())

	var bc api.BookController
	app.Get("/source", bc.Source)
	app.Get("/search", bc.Search)
	app.Get("/info", bc.Info)
	app.Get("/body", bc.Body)
	app.Post("/addBookSelf", bc.AddBookSelf)
	app.Get("/listBookSelf", bc.ListBookSelf)
	app.Post("/updateBookSelf", bc.UpdateBookSelf)

	return app
}

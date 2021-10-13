package main

import (
	_ "app/server/conf"
	"app/server/router"
	"log"
)

func main() {
	//app := router.New(fiber.Config{
	//	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
	//		if err != nil {
	//			return ctx.JSON(fiber.Map{
	//				"code":    201,
	//				"message": err.Error(),
	//			})
	//		}
	//		return nil
	//	},
	//})
	app := router.New()
	log.Fatal(app.Listen(":30080"))
}

package api

import (
	"app/server/common/result"
	"app/server/service"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
}

var bs service.BookService
var r result.Result

func (BookController) Search(ctx *fiber.Ctx) error {
	var s service.SearchReq
	if err := ctx.QueryParser(&s); err != nil {
		panic(err.Error())
	}
	sl := bs.Search(&s)
	return r.Ctx(ctx).SetData(sl).Success()
}

func (BookController) Info(ctx *fiber.Ctx) error {
	var i service.InfoReq
	if err := ctx.QueryParser(&i); err != nil {
		panic(err.Error())
	}
	b := bs.Info(&i)
	return r.Ctx(ctx).SetData(b).Success()
}

func (BookController) Body(ctx *fiber.Ctx) error {
	var b service.BodyReq
	if err := ctx.QueryParser(&b); err != nil {
		panic(err.Error())
	}
	body := bs.Body(&b)
	return r.Ctx(ctx).SetData(body).Success()
}

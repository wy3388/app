package api

import (
	"app/server/common/result"
	"app/server/model"
	"app/server/service"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
}

var bs service.BookService
var r result.Result

func (BookController) Source(ctx *fiber.Ctx) error {
	return r.Ctx(ctx).SetData(bs.Source()).Success()
}

func (BookController) Search(ctx *fiber.Ctx) error {
	var s service.SearchReq
	if err := ctx.QueryParser(&s); err != nil {
		panic(err)
	}
	sl := bs.Search(&s)
	return r.Ctx(ctx).SetData(sl).Success()
}

func (BookController) Info(ctx *fiber.Ctx) error {
	var i service.InfoReq
	if err := ctx.QueryParser(&i); err != nil {
		panic(err)
	}
	b := bs.Info(&i)
	return r.Ctx(ctx).SetData(b).Success()
}

func (BookController) Body(ctx *fiber.Ctx) error {
	var b service.BodyReq
	if err := ctx.QueryParser(&b); err != nil {
		panic(err)
	}
	body := bs.Body(&b)
	return r.Ctx(ctx).SetData(body).Success()
}

func (BookController) AddBookSelf(ctx *fiber.Ctx) error {
	var b model.BookSelf
	if err := ctx.BodyParser(&b); err != nil {
		panic(err)
	}
	bs.AddBookSelf(&b)
	return r.Ctx(ctx).Success()
}

func (BookController) ListBookSelf(ctx *fiber.Ctx) error {
	b := bs.ListBookSelf()
	return r.Ctx(ctx).SetData(b).Success()
}

func (BookController) UpdateBookSelf(ctx *fiber.Ctx) error {
	var req service.UpdateBookSelfReq
	if err := ctx.BodyParser(&req); err != nil {
		panic(err)
	}
	bs.UpdateBookSelf(&req)
	return r.Ctx(ctx).Success()
}

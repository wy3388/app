package result

import "github.com/gofiber/fiber/v2"

type ApiCode int

const (
	SUCCESS ApiCode = iota + 2000
	FAILED
)

type Result struct {
}

type ApiResult struct {
	Code    ApiCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var ctx *fiber.Ctx

func (Result) Ctx(c *fiber.Ctx) *ApiResult {
	ctx = c
	return &ApiResult{}
}

func (r *ApiResult) ApiResult(code ApiCode, message string, data interface{}) error {
	r.Code = code
	r.Message = message
	r.Data = data
	return ctx.JSON(r)
}

func (r *ApiResult) SetCode(code ApiCode) *ApiResult {
	r.Code = code
	return r
}

func (r *ApiResult) SetMessage(message string) *ApiResult {
	r.Message = message
	return r
}

func (r *ApiResult) SetData(data interface{}) *ApiResult {
	r.Data = data
	return r
}

func (r *ApiResult) Success() error {
	r.Code = SUCCESS
	if r.Message == "" {
		r.Message = "success"
	}
	if r.Data == nil {
		r.Data = map[string]string{}
	}
	return ctx.JSON(r)
}

func (r *ApiResult) Failed() error {
	r.Code = FAILED
	if r.Message == "" {
		r.Message = "failed"
	}
	if r.Data == nil {
		r.Data = map[string]string{}
	}
	return ctx.JSON(r)
}

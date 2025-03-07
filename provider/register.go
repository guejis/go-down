package provider

import (
	"down/helper"

	"github.com/gofiber/fiber/v2"
)

type RegisterComponent struct {
	Endpoint    string
	Method      string
	Title       string
	Description string
	Type 				string
	Params      map[string]interface{}
	Body        map[string]interface{}
	Code        func(*fiber.Ctx) error
}

type Register struct {
	Api []RegisterComponent
}

type UrlQuery struct {
	Url string `query:"url"`
}

type IDQuery struct {
	ID string `query:"id"`
}

var NewRegister *Register = &Register{}
var VS *helper.Visitor = &helper.Visitor{}

func (r *Register) RegisterProvider(i RegisterComponent) {
	r.Api = append(r.Api, i)
}

func (r *Register) GetRoutes() *Register {
	return r
}

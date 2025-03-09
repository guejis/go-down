package main

import (
	"fmt"
	"os"

	"down/provider"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type EndpointList struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Endpoint    string      `json:"endpoint"`
	Method      string      `json:"method"`
	Description string      `json:"description"`
	Params      interface{} `json:"params"`
	Body        interface{} `json:"body"`
	Type        string      `json:"type"`
	Hit         int         `json:"hit"`
	UpdateDate  string      `json:"date"`
	Status      string      `json:"status"`
	Demo        string      `json:"demo"`
}

var PORT string = os.Getenv("PORT")

func handleProvider(app *fiber.App) {
	prov := provider.NewRegister.GetRoutes()
	id := 1
	for i, v := range prov.Api {
		prov.Api[i].ID = id
		id++
		switch v.Method {
		case "GET":
			app.Get(v.Endpoint, v.Code)
		case "POST":
			app.Post(v.Endpoint, v.Code)
		case "PUT":
			app.Put(v.Endpoint, v.Code)
		}
	}
}

func useMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"resLen": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				str := fmt.Sprintf("%d bytes", len(c.Response().Body()))
				return output.WriteString(str)
			},
		},
		Format:        "[ ${time} ] ${ip} - ${status} - ${method} ${path} - ${resLen}${latency}\n",
		DisableColors: false,
		TimeFormat:    "2006-01-02T15:04:05-0700",
		TimeZone:      "Asia/Jakarta",
	}))
	app.Use(func(c *fiber.Ctx) error {
		var total any = provider.VS.Read(c.Path())
		if total == nil {
			total = 0
		}
		provider.VS.Write(c.Path(), total.(int)+1)
		return c.Next()
	})
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	useMiddleware(app)

	if PORT == "" {
		PORT = "3030"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://syntx-api.vercel.app")
	})

	app.Get("/list-endpoint", func(c *fiber.Ctx) error {
		var endpoint []EndpointList = make([]EndpointList, 0)
		list := provider.VS.ReadAll()

		for _, v := range provider.NewRegister.GetRoutes().Api {
			hit := 0

			for _, f := range list.Data {
				if f.Key == v.Endpoint {
					hit = f.Value.(int)
				}
			}

			demo := provider.BASE_API + v.Endpoint
			if v.Demo != "" {
				demo = v.Demo
			}

			endpoint = append(endpoint, EndpointList{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Endpoint:    v.Endpoint,
				Method:      v.Method,
				Params:      v.Params,
				Body:        v.Body,
				Type:        v.Type,
				Hit:         hit,
				UpdateDate:  "Sun 8:15pm",
				Status:      "Active",
				Demo:        demo,
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Hai 👋",
			"path":    endpoint,
		})
	})

	handleProvider(app)

	app.Listen(fmt.Sprintf(":%s", PORT))
}

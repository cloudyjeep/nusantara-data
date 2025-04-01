package api

import (
	"fmt"

	"github.com/cloudyjeep/nusantara-data/api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	*fiber.App
}

func (app *App) DefaultRoutes() {
	// main route
	app.Get("/", func(c *fiber.Ctx) error {
		h := config.LoadRequest(c)
		return h.ReturnMessage("Welcome to API Nusantara Data")
	})

	// handler 404 Not Found
	app.Use(func(c *fiber.Ctx) error {
		return config.LoadRequest(c).ReturnError(fiber.StatusNotFound, "Not Found", "The requested route does not exist")
	})
}

func (app *App) Middleware() {
	// logger
	app.Use(logger.New())

	// compression
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// access origin
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
}

func Init(port string) *App {
	app := App{fiber.New()}
	app.Middleware()

	AppRoutes(&app)
	app.DefaultRoutes()

	// run server api
	fmt.Printf("Server running at http://localhost:%v\n", port)
	app.Listen(fmt.Sprintf(":%v", port))
	return &app
}

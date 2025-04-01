package api

import (
	"github.com/cloudyjeep/nusantara-data/api/config"
	"github.com/cloudyjeep/nusantara-data/api/handler"
	"github.com/cloudyjeep/nusantara-data/api/service"
)

func AppRoutes(app *App) {
	// auth control
	auth := config.Auth{}

	// init service
	service := service.InitServices()

	// init handler
	category := handler.NewCategoryHandler(service)
	product := handler.NewProductHandler(service)

	// Endpoint public
	app.Get("/category", category.Find)
	app.Post("/category/:name", category.Create)
	app.Delete("/category/:name", category.Delete)

	// Endpoint protected
	app.Post("/product", auth.On(product.Create))
	app.Get("/product", auth.On(product.Find))
	app.Get("/product/:id", auth.On(product.FindById))
	app.Put("/product/:id", auth.On(product.Update))
	app.Delete("/product/:id", auth.On(product.Delete))
}

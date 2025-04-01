package handler

import (
	"github.com/cloudyjeep/nusantara-data/api/config"
	"github.com/cloudyjeep/nusantara-data/api/model"
	"github.com/cloudyjeep/nusantara-data/api/service"
	"github.com/gofiber/fiber/v2"
)

func NewProductHandler(instance *service.Injector) productHandler {
	return service.InjectService[product](instance)
}

type productHandler interface {
	// Handler > Get all product data
	Find(c *fiber.Ctx) error

	// Handler > Get product by id
	FindById(c *fiber.Ctx) error

	// Handler > Create new product
	Create(c *fiber.Ctx) error

	// Handler > Update product by id
	Update(c *fiber.Ctx) error

	// Handler > Delete product by id
	Delete(c *fiber.Ctx) error
}

type product service.Injector

// Create implements productHandler.
func (p product) Create(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	data := config.ReadBody[model.Product](c)
	return req.ReturnData(p.Service.Product.Create(data))
}

// Delete implements productHandler.
func (p product) Delete(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	return req.ReturnData(p.Service.Product.Delete(req.ParamId()))
}

// Find implements productHandler.
func (p product) Find(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	return req.ReturnData(p.Service.Product.FindByFilter(req.Pagination))
}

// FindById implements productHandler.
func (p product) FindById(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	return req.ReturnData(p.Service.Product.FindById(req.ParamId()))
}

// Update implements productHandler.
func (p product) Update(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	data := config.ReadBody[model.Product](c)
	data.Id = req.ParamId()
	return req.ReturnData(p.Service.Product.Update(data))
}

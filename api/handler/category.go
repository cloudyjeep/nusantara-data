package handler

import (
	"github.com/cloudyjeep/nusantara-data/api/config"
	"github.com/cloudyjeep/nusantara-data/api/model"
	"github.com/cloudyjeep/nusantara-data/api/service"
	"github.com/gofiber/fiber/v2"
)

func NewCategoryHandler(instance *service.Injector) categoryHandler {
	return service.InjectService[category](instance)
}

type categoryHandler interface {
	// Handler > Get data categories
	Find(c *fiber.Ctx) error

	// Handler > Create new category
	Create(c *fiber.Ctx) error

	// Handler > Delete category
	Delete(c *fiber.Ctx) error
}

type category service.Injector

// Create implements categoryHandler.
func (cat category) Create(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	data := model.Category(req.Param("name"))
	return req.ReturnData(cat.Service.Category.Create(data))

}

// Delete implements categoryHandler.
func (cat category) Delete(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	data := model.Category(req.Param("name"))
	return req.ReturnData(cat.Service.Category.Delete(data))
}

// Find implements categoryHandler.
func (cat category) Find(c *fiber.Ctx) error {
	req := config.LoadRequest(c)
	return req.ReturnData(cat.Service.Category.Find())
}

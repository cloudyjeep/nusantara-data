package config

import (
	"reflect"
	"strconv"

	"github.com/cloudyjeep/nusantara-data/lib"
	"github.com/gofiber/fiber/v2"
)

type Request func(c *fiber.Ctx) error

type request struct {
	ctx        *fiber.Ctx
	Pagination Pagination
}

// util handler

func (h *request) GetPagination() (page int, limit int) {
	page, _ = strconv.Atoi(h.ctx.Query("page"))
	limit, _ = strconv.Atoi(h.ctx.Query("limit"))
	return page, limit
}

// response

func (h *request) ReturnData(data any, err error) error {
	if err != nil {
		return h.ReturnDataError(err)
	}
	return h.ctx.JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})
}

func (h *request) ReturnDataPagination(data any, err error) error {
	if err != nil {
		return h.ReturnDataError(err)
	}
	return h.ctx.JSON(fiber.Map{
		"message": "success",
		"data":    data,
		"page":    h.Pagination.Page,
		"limit":   h.Pagination.Limit,
	})
}

func (h *request) ReturnMessage(message string) error {
	return h.ctx.JSON(fiber.Map{
		"message": message,
	})
}

func (h *request) ReturnError(status int, errors string, message any) error {
	h.ctx.Status(status)
	if reflect.ValueOf(message).IsValid() {
		return h.ctx.JSON(fiber.Map{
			"error":   errors,
			"message": message,
		})
	} else {
		return h.ctx.JSON(fiber.Map{
			"error": errors,
		})
	}
}

func (h *request) ParamId() string {
	return h.ctx.Params("id")
}

func (h *request) Param(name string) string {
	return h.ctx.Params(name)
}

func (h *request) ReturnDataError(err error) error {
	return h.ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": err.Error(),
	})
}

// load handler context
func LoadRequest(c *fiber.Ctx) *request {
	h := request{ctx: c}

	// validate pagination
	page, limit := h.GetPagination()
	h.Pagination.Page = lib.If(page > 0, page, 1)
	h.Pagination.Limit = lib.If(limit > 0, limit, 10)

	return &h
}

func ReadBody[T any](c *fiber.Ctx) T {
	var ref T
	c.BodyParser(&ref)
	return ref
}

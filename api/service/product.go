package service

import (
	"fmt"
	"slices"
	"time"

	"github.com/cloudyjeep/nusantara-data/api/config"
	"github.com/cloudyjeep/nusantara-data/api/model"
	"github.com/cloudyjeep/nusantara-data/lib"
	"github.com/google/uuid"
)

func NewProductService(instance *Injector) productService {
	return InjectService[product](instance)
}

type productService interface {
	// Service > Find product by id
	FindById(id string) (*model.Product, error)

	// Service > Find product by filter pagination
	FindByFilter(filter config.Pagination) ([]model.Product, error)

	// Service > Create new product
	Create(data model.Product) (*model.Product, error)

	// Service > Update product
	Update(data model.Product) (*model.Product, error)

	// Service > Delete product
	Delete(id string) (*model.Product, error)
}

var dummyProductService = model.GenerateProductDummy()

type product Injector

// Create implements productService.
func (p product) Create(data model.Product) (*model.Product, error) {
	data.Name = lib.Trim(data.Name)

	if data.Name == "" {
		return nil, fmt.Errorf("name can't empty")
	}

	today := time.Now()
	data.Id = uuid.NewString()
	data.CreatedAt = today
	data.UpdatedAt = today
	dummyProductService = append(dummyProductService, data)

	return &data, nil
}

// Delete implements productService.
func (p product) Delete(id string) (*model.Product, error) {
	item, idx := p.getItemById(id)

	if item != nil {
		dummyProductService = slices.Delete(dummyProductService, idx, idx+1)
		return item, nil
	}

	return nil, fmt.Errorf("data not found")
}

// FindByFilter implements productService.
func (p product) FindByFilter(filter config.Pagination) ([]model.Product, error) {
	start, end := lib.OffsetStartEnd(filter.Page, filter.Limit)
	fmt.Printf("start: %v - end: %v\n", start, end)
	return lib.CutSlices(dummyProductService, start, end), nil
}

// FindById implements productService.
func (p product) FindById(id string) (*model.Product, error) {
	item, _ := p.getItemById(id)

	if item != nil {
		return item, nil
	}

	return nil, fmt.Errorf("data not found")
}

// Update implements productService.
func (p product) Update(data model.Product) (*model.Product, error) {
	data.Id = lib.Trim(data.Id)
	data.Name = lib.Trim(data.Name)

	if data.Id == "" || data.Name == "" {
		return nil, fmt.Errorf("id or name can't empty")
	}

	item, idx := p.getItemById(data.Id)
	if item != nil {
		dummyProductService[idx].Name = data.Name
		dummyProductService[idx].Price = data.Price
		dummyProductService[idx].Stock = data.Stock
		dummyProductService[idx].UpdatedAt = time.Now()
		return &data, nil
	}

	return nil, fmt.Errorf("data not found")
}

// internal
func (p product) getItemById(id string) (item *model.Product, idx int) {
	count := len(dummyProductService)

	for i := 0; i < count; i++ {
		item := dummyProductService[i]
		if item.Id == id {
			return &item, i
		}
	}

	return nil, 0
}

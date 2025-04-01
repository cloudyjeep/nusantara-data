package service

import (
	"sort"

	"github.com/cloudyjeep/nusantara-data/api/model"
)

func NewCategoryService(instance *Injector) categoryService {
	return InjectService[category](instance)

}

type categoryService interface {
	// Service > Find category
	Find() ([]model.Category, error)

	// Service > Create new category
	Create(data model.Category) (result []model.Category, err error)

	// Service > Delete category
	Delete(name model.Category) (result []model.Category, err error)
}

var dummyCategoryService = model.GenerateCategoryDummy()

type category Injector

// Create implements categoryService.
func (c category) Create(data model.Category) ([]model.Category, error) {
	c.save(append(dummyCategoryService, data))
	return c.Find()
}

// Delete implements categoryService.
func (c category) Delete(name model.Category) (result []model.Category, err error) {
	count := len(dummyCategoryService)
	for i := 0; i < count; i++ {
		if dummyCategoryService[i] != name {
			result = append(result, dummyCategoryService[i])
		}
	}
	c.save(result)
	return
}

// Find implements categoryService.
func (c category) Find() ([]model.Category, error) {
	return dummyCategoryService, nil
}

// method
func (c category) save(values []model.Category) {
	// sort category
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	dummyCategoryService = values
}

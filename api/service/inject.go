package service

import "reflect"

type Injector struct {
	Service *service
}

type service struct {
	// declare instance service here
	Product  productService
	Category categoryService
}

// init and register service
func InitServices() *Injector {
	instance := Injector{Service: &service{}}

	// register service here
	instance.Service.Product = NewProductService(&instance)
	instance.Service.Category = NewCategoryService(&instance)

	return &instance
}

func InjectService[T any](instance any) T {
	var ref T
	reflect.ValueOf(&ref).Elem().FieldByName("Service").Set(
		reflect.ValueOf(instance).Elem().FieldByName("Service"),
	)
	return ref
}

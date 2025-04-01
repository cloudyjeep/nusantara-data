package model

import (
	"time"

	"github.com/cloudyjeep/nusantara-data/lib"
)

type Product struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     float64   `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GenerateProductDummy() (result []Product) {
	result, _ = lib.ReadJSONFile[[]Product]("assets/product-dummy.json")
	return result
}

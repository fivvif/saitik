package main

import (
	product3 "saitik/internal/adapters/api/product"
	product2 "saitik/internal/adapters/db/product"
	"saitik/internal/domain/product"
)

func main() {
	productStorage := product2.NewStorage()

	productService := product.NewService(productStorage)

	productHandler := product3.NewHandler(productService)

}

package product

import (
	"context"
	"saitik/internal/domain/product"
)

type Service interface {
	GetProdByUUID(ctx context.Context, uuid string) *product.Product
	CreateProduct(ctx context.Context, dto *product.CreateProductDTO) *product.Product
}

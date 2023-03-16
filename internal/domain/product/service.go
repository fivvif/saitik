package product

import (
	"context"
	"saitik/internal/adapters/api/product"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) product.Service {
	return &service{storage: storage}

}

func (s *service) GetProdByUUID(ctx context.Context, uuid string) *Product {
	return s.storage.GetOne(uuid)
}

func (s *service) CreateProduct(ctx context.Context, dto *CreateProductDTO) *Product {

}

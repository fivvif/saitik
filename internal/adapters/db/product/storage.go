package product

import "saitik/internal/domain/product"

type productStorage struct {
}

func NewStorage() product.Storage {
	return &productStorage{}

}

func (ps *productStorage) GetOne(uuid string) *product.Product {
	return nil
}
func (ps *productStorage) Create(product *product.Product) *product.Product {
	return nil
}
func (ps *productStorage) Delete(product *product.Product) error {
	return nil
}

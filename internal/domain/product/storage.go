package product

type Storage interface {
	GetOne(uuid string) *Product
	Create(product *Product) *Product
	Delete(product *Product) error
}

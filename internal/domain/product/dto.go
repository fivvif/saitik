package product

type CreateProductDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type UpdateProductDTO struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	LastUpdated string `json:"last_updated"`
}

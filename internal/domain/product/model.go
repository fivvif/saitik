package product

type Product struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Price       string `json:"price,omitempty"`
	LastUpdated string `json:"last_updated,omitempty"`
}

func (p *Product) Find(name string) {

}

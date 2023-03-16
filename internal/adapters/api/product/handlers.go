package product

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"saitik/internal/adapters/api"
)

const (
	productURL  = "/products/:product_id"
	productsURL = "/products"
)

type handler struct {
	productService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{productService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(productsURL, h.FindProd)
}

func (h *handler) FindProd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// product := h.productService.GetProdByUUID(context.Background(), 1)
	w.Write([]byte("prodname"))
	w.WriteHeader(http.StatusOK)

}

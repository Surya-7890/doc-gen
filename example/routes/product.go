package routes

import "net/http"

type ProductRouter struct {
	Mux *http.ServeMux
}

func NewProductRouter(mux *http.ServeMux) *ProductRouter {
	return &ProductRouter{
		Mux: mux,
	}
}

func (p *ProductRouter) GetAllProducts() {}

func (p *ProductRouter) GetProductById() {}

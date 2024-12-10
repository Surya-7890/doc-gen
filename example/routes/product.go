package routes

import (
	"fmt"
	"net/http"
)

type ProductRouter struct {
	Mux *http.ServeMux
}

func NewProductRouter(mux *http.ServeMux) *ProductRouter {
	return &ProductRouter{
		Mux: mux,
	}
}

func (p *ProductRouter) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("no body here to parse")
}

func (p *ProductRouter) GetProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("no body here to parse, but the route path should be parsed for getting id")
}

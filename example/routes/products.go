package routes

import (
	"gen-doc/example/controllers"
	"gen-doc/example/db"
	"net/http"
)

type ProductRoute struct {
	Handler *http.ServeMux
}

func NewProductRoute(db *db.DB) *ProductRoute {
	return &ProductRoute{
		Handler: setupProductHandler(db),
	}
}

func setupProductHandler(db *db.DB) *http.ServeMux {
	mux := http.NewServeMux()
	product_controller := controllers.NewProductController(db)

	mux.HandleFunc("GET /", product_controller.GetMethod)
	mux.HandleFunc("POST /", product_controller.PostMethod)
	mux.HandleFunc("PATCH /", product_controller.PatchMethod)
	mux.HandleFunc("DELETE /", product_controller.DeleteMethod)

	return mux
}

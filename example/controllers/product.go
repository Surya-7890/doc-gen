package controllers

import (
	"gen-doc/example/db"
	"net/http"
)

type ProductController struct {
	DB *db.DB
}

func NewProductController(db *db.DB) *ProductController {
	return &ProductController{
		DB: db,
	}
}

func (a *ProductController) GetMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *ProductController) PostMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *ProductController) PatchMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *ProductController) DeleteMethod(w http.ResponseWriter, r *http.Request) {

}

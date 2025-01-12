package controllers

import (
	"encoding/json"
	"gen-doc/example/db"
	"gen-doc/example/types"
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

// @GET /product
// @desc get product list
func (a *ProductController) GetMethod(w http.ResponseWriter, r *http.Request) {
	var res *types.SampleResponse2
	json.NewEncoder(w).Encode(res)
}

// @POST /product
func (a *ProductController) PostMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest2

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse2
	json.NewEncoder(w).Encode(res)
}

// @PATCH /product
func (a *ProductController) PatchMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest2

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse2
	json.NewEncoder(w).Encode(res)
}

// @DELETE /product
func (a *ProductController) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest2

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse2
	json.NewEncoder(w).Encode(res)
}

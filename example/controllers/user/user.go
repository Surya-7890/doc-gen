package controllers

import (
	"encoding/json"
	"gen-doc/example/db"
	"gen-doc/example/types"
	"net/http"
)

type UserController struct {
	DB *db.DB
}

func NewUserController(db *db.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

// @GET /user/{id}
func (a *UserController) GetMethod(w http.ResponseWriter, r *http.Request) {
	var res *types.SampleResponse1
	json.NewEncoder(w).Encode(res)
}

// @POST /user
func (a *UserController) PostMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest1

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse1
	json.NewEncoder(w).Encode(res)
}

// @PATCH /user
func (a *UserController) PatchMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest1

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse1
	json.NewEncoder(w).Encode(res)
}

// @DELETE /user
func (a *UserController) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest1

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse1
	json.NewEncoder(w).Encode(res)
}

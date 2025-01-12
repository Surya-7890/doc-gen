package controllers

import (
	"encoding/json"
	"gen-doc/example/db"
	"gen-doc/example/types"
	"net/http"
)

type UserProfileController struct {
	DB *db.DB
}

func NewUserProfileController(db *db.DB) *UserProfileController {
	return &UserProfileController{
		DB: db,
	}
}

// @GET /user/product
func (a *UserProfileController) GetMethod(w http.ResponseWriter, r *http.Request) {
	var res *types.SampleResponse
	json.NewEncoder(w).Encode(res)
}

// @POST /user/product
func (a *UserProfileController) PostMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse
	json.NewEncoder(w).Encode(res)
}

// @PATCH /user/product
func (a *UserProfileController) PatchMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse
	json.NewEncoder(w).Encode(res)
}

// @DELETE /user/product
func (a *UserProfileController) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	var req *types.SampleRequest

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}

	var res *types.SampleResponse
	json.NewEncoder(w).Encode(res)
}

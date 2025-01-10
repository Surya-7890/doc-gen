package controllers

import (
	"gen-doc/example/db"
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

}

// @POST /user/product
func (a *UserProfileController) PostMethod(w http.ResponseWriter, r *http.Request) {

}

// @PATCH /user/product
func (a *UserProfileController) PatchMethod(w http.ResponseWriter, r *http.Request) {

}

// @DELETE /user/product
func (a *UserProfileController) DeleteMethod(w http.ResponseWriter, r *http.Request) {

}

package controllers

import (
	"gen-doc/example/db"
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

// @GET /user
func (a *UserController) GetMethod(w http.ResponseWriter, r *http.Request) {

}

// @POST /user
func (a *UserController) PostMethod(w http.ResponseWriter, r *http.Request) {

}

// @PATCH /user
func (a *UserController) PatchMethod(w http.ResponseWriter, r *http.Request) {

}

// @DELETE /user
func (a *UserController) DeleteMethod(w http.ResponseWriter, r *http.Request) {

}

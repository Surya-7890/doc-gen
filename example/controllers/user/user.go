package user

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

func (a *UserController) GetMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserController) PostMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserController) PatchMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserController) DeleteMethod(w http.ResponseWriter, r *http.Request) {

}

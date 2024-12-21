package user

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

func (a *UserProfileController) GetMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserProfileController) PostMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserProfileController) PatchMethod(w http.ResponseWriter, r *http.Request) {

}

func (a *UserProfileController) DeleteMethod(w http.ResponseWriter, r *http.Request) {

}

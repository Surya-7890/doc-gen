package routes

import (
	controllers "gen-doc/example/controllers/user"
	"gen-doc/example/db"
	"net/http"
)

type UserProfileRoute struct {
	Handler *http.ServeMux
}

func NewUserProfileRoute(db *db.DB) *UserProfileRoute {
	return &UserProfileRoute{
		Handler: setupUserProfileHandler(db),
	}
}

func setupUserProfileHandler(db *db.DB) *http.ServeMux {
	mux := http.NewServeMux()
	user_profile_controller := controllers.NewUserProfileController(db)

	mux.HandleFunc("GET /", user_profile_controller.GetMethod)
	mux.HandleFunc("POST /", user_profile_controller.PostMethod)
	mux.HandleFunc("PATCH /", user_profile_controller.PatchMethod)
	mux.HandleFunc("DELETE /", user_profile_controller.DeleteMethod)

	return mux
}

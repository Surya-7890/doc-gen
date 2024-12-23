package routes

import (
	controllers "gen-doc/example/controllers/user"
	"gen-doc/example/db"
	"net/http"
)

type UserRoute struct {
	Handler *http.ServeMux
}

func NewUserRoute(db *db.DB) *UserRoute {
	return &UserRoute{
		Handler: setupUserHandler(db),
	}
}

func setupUserHandler(db *db.DB) *http.ServeMux {
	mux := http.NewServeMux()
	profile := NewUserProfileRoute(db)
	user_controller := controllers.NewUserController(db)
	mux.Handle("/profile", http.StripPrefix("/profile", profile.Handler))

	mux.HandleFunc("GET /", user_controller.GetMethod)
	mux.HandleFunc("POST /", user_controller.PostMethod)
	mux.HandleFunc("PATCH /", user_controller.PatchMethod)
	mux.HandleFunc("DELETE /", user_controller.DeleteMethod)

	return mux
}

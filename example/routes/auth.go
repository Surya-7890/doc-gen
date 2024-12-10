package routes

import "net/http"

type AuthRouter struct {
	Mux *http.ServeMux
}

func NewAuthRouter(mux *http.ServeMux) *AuthRouter {
	return &AuthRouter{
		Mux: mux,
	}
}

func (a *AuthRouter) Login() {}

func (a *AuthRouter) Signup() {}

package routes

import (
	"encoding/json"
	"fmt"
	"gen-doc/example/types"
	"log"
	"net/http"
)

type AuthRouter struct {
	Mux *http.ServeMux
}

func NewAuthRouter(mux *http.ServeMux) *AuthRouter {
	return &AuthRouter{
		Mux: mux,
	}
}

func (a *AuthRouter) Login(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(user)
}

func (a *AuthRouter) Signup(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(user)
}

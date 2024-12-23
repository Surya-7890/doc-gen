package main

import (
	"gen-doc/example/db"
	"gen-doc/example/routes"
	user "gen-doc/example/routes/user"
	"net/http"
)

func main() {
	mux := setupApplication()

	server := http.Server{
		Addr:    ":7000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func connectToDB() *db.DB {
	return db.NewDB()
}

func setupApplication() *http.ServeMux {

	mux := http.NewServeMux()
	db := connectToDB()

	user := user.NewUserRoute(db)
	product := routes.NewProductRoute(db)

	mux.Handle("/", http.StripPrefix("/product", product.Handler))
	mux.Handle("/", http.StripPrefix("/user", user.Handler))

	return mux
}

package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	admin := http.NewServeMux()
	mux.Handle("/admin", admin)
}

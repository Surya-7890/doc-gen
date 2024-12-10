package main

import (
	"gen-doc/parser"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	admin := http.NewServeMux()
	mux.Handle("/admin", admin)
	logger := log.Default()

	Parser := parser.NewParser(logger)

	Parser.Parse()

}

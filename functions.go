package main

import (
	"fmt"
	"net/http"
)

type Response struct{}

type Request struct{}

func Hello(res http.ResponseWriter, req Request) {
	fmt.Println(req, res)
}

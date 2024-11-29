package main

import "fmt"

type Response struct{}

type Request struct{}

func Hello(res Response, req *Request) {
	fmt.Println(req, res)
}

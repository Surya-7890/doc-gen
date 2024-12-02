package test

import (
	"net/http"
)

func IncorrectParams(res string, req int) {}

func IncorrectFirstParam(res string, req *http.Request) {}

func IncorrectNoOfParams(res string) {}

func IncorrectSecondParam(res http.ResponseWriter, req string) {}

// @method POST
// @path /login
func RouteHandler(res http.ResponseWriter, req *http.Request) {}

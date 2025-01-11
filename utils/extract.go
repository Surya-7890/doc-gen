package utils

import (
	"errors"
	"fmt"
	"gen-doc/types"
	"strings"
)

func ExtractDocFromText(doc types.ROUTE_INFO, res *types.RouteInfo) {

	res.RequestParsingRequired = true

	req_type, ok := doc.Identifier.(types.REQUEST_TYPE)
	if ok {
		res.Method = strings.TrimPrefix(req_type.ToString(), "@")
		if req_type == types.GET || req_type == types.OPTIONS {
			res.RequestParsingRequired = false
		}
		if req_type == "@desc" {
			fmt.Println(doc.Identifier)
		}
		res.Route = doc.Value
		params, err := getPathParams(res.Route)
		if err != nil {
			res = nil
		}
		res.PathParams = params
		return
	}

	_, ok = doc.Identifier.(types.DESCRIPTION)
	if ok {
		res.Description = doc.Value
	}

}

func getPathParams(route string) ([]string, error) {
	params := []string{}
	length := len(route)

	for i := 0; i < length; i++ {

		// start of a new path param
		if route[i] == '{' {

			found := false // to check for invalid syntax
			curr := ""
			for j := i + 1; j < length; j++ {
				if route[j] == '}' {
					found = true
					break
				}
				curr += string(route[j])
			}

			if !found {
				return params, errors.New("invalid syntax")
			}

			params = append(params, curr)
		}
	}

	return params, nil
}

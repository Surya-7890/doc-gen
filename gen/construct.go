package gen

import (
	"fmt"
	"gen-doc/types"
	"go/ast"
	"strings"
)

func (g *Gen) constructRouteInfo(fn *ast.FuncDecl) {
	route := types.Route{}
	for _, val := range fn.Doc.List {
		str := strings.TrimPrefix(val.Text, "// ")
		arr := strings.Split(str, " ")
		if len(arr) != 2 {
			fmt.Println("not enough arguements for " + fn.Name.Name)
		}

		if arr[0] == string(types.METHOD) {
			route.Method = types.HTTP_METHOD(arr[1])
		}
		if arr[0] == string(types.PATH) {
			route.Path = arr[1]
		}
	}

	val, ok := types.HasBodyMethods[route.Method]
	if ok && val {
		g.constructRequestBody(fn)
	} else {
		route.Body = nil
	}

	fmt.Println(route)
}

func (g *Gen) constructRequestBody(fn *ast.FuncDecl) {
	fmt.Println(fn.Body)
}

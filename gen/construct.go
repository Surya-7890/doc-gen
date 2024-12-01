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

		if arr[0] == types.METHOD_ {
			route.Method = arr[1]
		}
		if arr[0] == types.PATH_ {
			route.Path = arr[1]
		}

		g.generateSwaggerJSON()
	}
	fmt.Println(route)
}

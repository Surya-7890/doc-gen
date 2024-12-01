package gen

import (
	"fmt"
	"go/ast"
	"strings"
)

func (g *Gen) constructRouteInfo(fn *ast.FuncDecl) {
	fmt.Println(fn.Name.Name)
	for _, val := range fn.Doc.List {
		str := strings.TrimPrefix(val.Text, "// ")
		arr := strings.Split(str, " ")
		fmt.Println(arr)
	}
}

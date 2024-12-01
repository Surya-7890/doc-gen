package parser

import (
	"fmt"
	"go/ast"
)

func (p *Parser) Traverse(file *ast.File) []*ast.FuncDecl {
	arr := []*ast.FuncDecl{}
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if !p.IsRouteHandler(fn) {
			messaage := fmt.Sprintf("function %s is not a route handler", fn.Name.Name)
			p.Log.Println(messaage)
			return true
		}

		message := fmt.Sprintf("%s is a route handler", fn.Name.Name)
		p.Log.Println(message)
		if fn != nil {
			arr = append(arr, fn)
		}

		return true
	})

	return arr
}

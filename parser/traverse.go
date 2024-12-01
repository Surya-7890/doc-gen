package parser

import (
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
			return true
		}

		if fn != nil {
			arr = append(arr, fn)
		}

		return true
	})

	return arr
}

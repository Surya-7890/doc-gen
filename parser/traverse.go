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

		fn_call, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		ident, ok := fn_call.X.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Name != "http" || fn_call.Sel.Name != "HandleFunc" {
			return true
		}

		fmt.Println(fn_call.X, fn_call.Sel)

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

func (p *Parser) findHandlerFunc() {

}

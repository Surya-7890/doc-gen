package parser

import (
	"fmt"
	"go/ast"
)

// checks if a function is a route handler
// reads the params to verify
// route handlers look like: HandlerName(res http.ResponseWriter, req *http.Request)
func (p *Parser) IsRouteHandler(fn *ast.FuncDecl) bool {

	params := fn.Type.Params.List
	if len(params) != 2 {
		return false
	}

	switch t := params[0].Type.(type) {
	case *ast.SelectorExpr:
		x, ok := t.X.(*ast.Ident)
		if !ok || x.Name != "http" || t.Sel.Name != "ResponseWriter" {
			return false
		}
	default:
		return false
	}

	switch t := params[1].Type.(type) {
	case *ast.StarExpr:
		val, ok := t.X.(*ast.SelectorExpr)
		if !ok {
			return false
		}

		if x, ok := val.X.(*ast.Ident); ok && x.Name == "http" && val.Sel.Name == "Request" {
			return true
		}

		return false
	default:
		return false
	}
}

func (p *Parser) IsHandlerFunc(fn *ast.FuncDecl) bool {
	fmt.Println(fn.Name)
	return true
}

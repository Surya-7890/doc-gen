package parser

import (
	"fmt"
	"gen-doc/types"
	"go/ast"
)

func (p *Parser) Traverse(file *ast.File) []*ast.FuncDecl {
	arr := []*ast.FuncDecl{}
	handlers := []ast.Node{}
	handlerFuncs := []ast.Node{}

	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		switch t := p.findFuncs(n, handlers, handlerFuncs); t {
		case types.HANDLE:
		case types.HANDLE_FUNC:
		case types.NONE:
		default:
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

func (p *Parser) findFuncs(n ast.Node, handlers, handlerFuncs []ast.Node) types.FUNC_TYPE {
	fn_call, ok := n.(*ast.SelectorExpr)
	if !ok {
		return types.NONE
	}

	ident, ok := fn_call.X.(*ast.Ident)
	if !ok {
		return types.NONE
	}

	if (ident.Name == "http" || ident.Name == "mux") && (fn_call.Sel.Name == "HandleFunc") {
		handlerFuncs = append(handlerFuncs, n)
		return types.HANDLE_FUNC
	}

	if (ident.Name == "http" || ident.Name == "mux") && (fn_call.Sel.Name == "Handle") {
		handlers = append(handlers, n)
		return types.HANDLE
	}

	return types.NONE
}

func (p *Parser) findHandlerFunc(n ast.Node) ast.Node {
	fn, ok := n.(*ast.FuncDecl)
	if !ok {
		return nil
	}
	param := fn.Type.Params.List

	if len(param) == 2 {
		// if p.checkIfHandlerFunc(n) {
		// 	return n
		// }
	}
	body := fn.Body.List

	for _, val := range body {
		select_stmt, ok := val.(*ast.SelectStmt)
		if !ok {
			continue
		}
		fmt.Println(select_stmt)
	}

	return nil
}

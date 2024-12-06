package scanner

import (
	"fmt"
	"go/ast"
	"go/token"
)

// find all function calls for http.Handle or mux.Handle
func (s *Scanner) findHandlers(n ast.Node) {
	call_expr, ok := n.(*ast.CallExpr)
	if !ok {
		return
	}

	stmt, ok := call_expr.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	stmt_ident, ok := stmt.X.(*ast.Ident)
	if !ok {
		return
	}

	if (stmt_ident.Name == "http" || stmt_ident.Name == "mux") && stmt.Sel.Name == "Handle" {
		s.findHandlerFuncs(call_expr)
	}
}

// find all function calls for http.HandleFunc or mux.HandleFunc
func (s *Scanner) findHandlerFuncs(n *ast.CallExpr) {
	if len(n.Args) != 2 {
		return
	}

	literal, ok := n.Args[0].(*ast.BasicLit)
	if !ok {
		return
	}

	if literal.Kind != token.STRING {
		return
	}
	fmt.Println(n.Args[1])
}

package scanner

import (
	"go/ast"
)

// find all function calls for http.Handle or mux.Handle
func (s *Scanner) findHandlers(node *ast.FuncDecl) []ast.Expr {
	var found []ast.Expr

	ast.Inspect(node, func(n ast.Node) bool {
		call_expr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		stmt, ok := call_expr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		stmt_ident, ok := stmt.X.(*ast.Ident)
		if !ok {
			return true
		}

		if (stmt_ident.Name == "http" || stmt_ident.Name == "mux") && stmt.Sel.Name == "Handle" {
			found = call_expr.Args
			return false
		}

		return true
	})
	return found
}

// // find all function calls for http.HandleFunc or mux.HandleFunc
// func (s *Scanner) findHandlerFuncs(n *ast.CallExpr) {
// 	if len(n.Args) != 2 {
// 		return
// 	}

// 	literal, ok := n.Args[0].(*ast.BasicLit)
// 	if !ok {
// 		return
// 	}

// 	if literal.Kind != token.STRING {
// 		return
// 	}
// 	// find where the
// 	fmt.Println(n.Args[1])
// }

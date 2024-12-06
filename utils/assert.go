package utils

import (
	"gen-doc/types"
	"go/ast"
)

func Assert(n interface{}, node_type types.AST_TYPE) bool {
	switch node_type {
	case types.CALL_EXPR:
		_, ok := n.(*ast.CallExpr)
		return ok
	case types.EXPR_STMT:
		_, ok := n.(*ast.ExprStmt)
		return ok
	default:
		return false
	}
}

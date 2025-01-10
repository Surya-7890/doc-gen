package utils

import "go/ast"

func CheckIfInBuiltFunction[T *ast.CallExpr | *ast.SelectorExpr](fn T) {
	switch any(fn).(type) {
	case *ast.CallExpr:
	case *ast.SelectorExpr:
	default:
	}
}

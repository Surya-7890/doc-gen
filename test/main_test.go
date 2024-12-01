package test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var (
	Map = make(map[string]*ast.FuncDecl)
)

func TestMain(m *testing.M) {
	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, "./functions.go", nil, parser.AllErrors)
	if err != nil {
		panic(err.Error())
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		if fn, ok := n.(*ast.FuncDecl); ok {
			Map[fn.Name.Name] = fn
			return true
		}

		return true
	})

	m.Run()
}

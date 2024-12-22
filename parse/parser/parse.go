package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func (p *Parser) parseSingleFile(filaname string, fileSet *token.FileSet, mode parser.Mode) ([]*ast.FuncDecl, []*ast.CallExpr) {
	fn := []*ast.FuncDecl{}
	call := []*ast.CallExpr{}

	file, err := parser.ParseFile(fileSet, filaname, nil, mode)
	if err != nil {
		p.log.Fatal(err.Error())
	}

	ast.Inspect(file, func(n ast.Node) bool {

		fn_decl, ok := n.(*ast.FuncDecl)
		if ok {
			fn = append(fn, fn_decl)
			return true
		}

		call_expr, ok := n.(*ast.CallExpr)
		if ok {
			// have to check if it is a builtin function
			// or it is from the std library
			call = append(call, call_expr)
			return true
		}

		return true
	})

	return fn, call
}

package parser

import (
	"go/ast"
)

func (p *Parser) parseSingleFile(file *ast.File) []*ast.FuncDecl {
	fn := []*ast.FuncDecl{}

	ast.Inspect(file, func(n ast.Node) bool {

		fn_decl, ok := n.(*ast.FuncDecl)
		if ok {
			fn = append(fn, fn_decl)
			return true
		}

		return true
	})

	return fn
}

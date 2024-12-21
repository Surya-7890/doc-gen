package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func (p *Parser) parseSingleFile(filaname string, fileSet *token.FileSet, mode parser.Mode) {
	file, err := parser.ParseFile(fileSet, filaname, nil, mode)
	if err != nil {
		p.log.Fatal(err.Error())
	}

	ast.Inspect(file, func(n ast.Node) bool {
		return true
	})
}

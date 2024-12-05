package parser

import (
	"fmt"
	"go/ast"
)

func (p *Parser) traverseHandler(n ast.Node) {}

func (p *Parser) traverseHandlerFunc(n ast.Node) {
	fn, ok := n.(*ast.FuncDecl)
	if !ok {
		return
	}
	for _, val := range fn.Body.List {
		// select statement
		// this means that this current function
		// is not the HandlerFunc itself
		if slct, ok := val.(*ast.SelectStmt); ok {
			fmt.Println(slct.Body)
			continue
		}

		if expr_stmt, ok := val.(*ast.ExprStmt); ok {
			if fn_call, ok := expr_stmt.X.(*ast.SelectorExpr); ok {

				ident, ok := fn_call.X.(*ast.Ident)
				if !ok {
					continue
				}

				// handle different types of parsing here
				// 1) json.NewDecoder().Decode() ~ also for xml
				// 1.1) decoder := json.NewDecoder()
				//      decoder.Decode()
				// 2) io.ReadAll()
				// 3) bufio.NewScanner().Scan() ~ for streams
				// 3.1) scanner := bufio.NewScanner()
				//      for scanner.Scan() {}
				// 4) io.TeeReader() ~ io.ReadAll after this

				fmt.Println(ident.Name)
			}
		}

	}
}

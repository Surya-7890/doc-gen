package parser

import (
	"fmt"
	"go/ast"
)

// traverses from the main function through routers (*http.ServeMux)
// and through routers from other packages till a HandlerFunc is reached
func (p *Parser) traverse() {
	for _, val := range p.fn_decls["main"] {
		if val.Name.Name == "main" {
			p.traverseFunc("main", "main", val)
		}
	}
}

// todo ~ traverse call_expr map to find all callExpr inside function body
// then params are checked if it is a http.ServeMux call expr
func (p *Parser) traverseFunc(pkg_name, fn_name string, fn *ast.FuncDecl) {
	for _, v := range fn.Body.List {
		assign, ok := v.(*ast.AssignStmt)
		if !ok {
			continue
		}

		for _, lhs := range assign.Lhs {
			fmt.Println(lhs)
		}

	}
	p.log.Println("len: ", len(fn.Body.List))
}

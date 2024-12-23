package parser

import (
	"fmt"
	"gen-doc/utils"
	"go/ast"
)

func (p *Parser) filterHandlerFuncs() {
	filtered := utils.Filter(p.fn_decls, func(entry *ast.FuncDecl) bool {
		params := entry.Type.Params.List
		if len(params) != 2 {
			return false
		}

		sl_res, ok := params[0].Type.(*ast.SelectorExpr)
		if !ok || sl_res.Sel.Name != "ResponseWriter" {
			return false
		}
		id_res, ok := sl_res.X.(*ast.Ident)
		if !ok || id_res.Name != "http" {
			return false
		}

		// fmt.Println(id_res)
		star_exp, ok := params[1].Type.(*ast.StarExpr)
		if !ok {
			return false
		}

		sl_req, ok := star_exp.X.(*ast.SelectorExpr)
		if !ok || sl_req.Sel.Name != "Request" {
			return false
		}
		id_req, ok := sl_req.X.(*ast.Ident)
		if !ok || id_req.Name != "http" {
			return false
		}

		fmt.Println(id_res.Name, sl_res.Sel.Name, id_req.Name, sl_req.Sel.Name)

		return true
	})

	// contains only handler funcs
	fmt.Println(len(filtered))
}

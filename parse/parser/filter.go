package parser

import (
	"gen-doc/utils"
	"go/ast"
)

func (p *Parser) filterHandlerFuncs() {
	for k := range p.fn_decls {
		filtered := utils.Filter(p.fn_decls[k], func(entry *ast.FuncDecl) bool {
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

			return true
		})

		// contains only handler funcs
		// maps [package_name] -> []HandlerFuncs (handler functions in the package)
		p.handler_funcs[k] = filtered
	}

}

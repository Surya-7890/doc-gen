package parser

import (
	"fmt"
	"gen-doc/types"
	"gen-doc/utils"
	"go/ast"
	"strings"
)

// function used to extract all @ symbols
// used to represent a route
func (p *Parser) extractDocs() {
	for pkgname, fns := range p.handler_funcs {
		p.log.Println(pkgname)
		for _, fn := range fns {
			res := &types.RouteInfo{}

			for _, doc := range fn.Doc.List {
				doc.Text = strings.TrimPrefix(doc.Text, "//")
				doc.Text = strings.TrimSpace(doc.Text)

				docArr := strings.Fields(doc.Text)
				if len(docArr) < 2 {
					continue
				}

				// Determine whether it's a REQUEST_TYPE or DESCRIPTION
				var identifier types.IDENTIFIER
				switch strings.ToLower(docArr[0]) {
				case "@get", "@post", "@patch", "@put", "@delete", "@options":
					identifier = types.REQUEST_TYPE(strings.ToLower(docArr[0]))
				default:
					identifier = types.DESCRIPTION(docArr[0])
				}

				utils.ExtractDocFromText(types.ROUTE_INFO{
					Identifier: identifier,
					Value:      strings.Join(docArr[1:], " "),
				}, res)
			}

			if res.RequestParsingRequired {
				p.parseStruct(fn, types.PARSE_REQUEST)
			}

			p.parseStruct(fn, types.PARSE_RESPONSE)
		}
	}
}

func (p *Parser) parseStruct(fn *ast.FuncDecl, target int) {
	if target == types.PARSE_REQUEST {

	} else if target == types.PARSE_RESPONSE {

	}

	for _, stmt := range fn.Body.List {
		var expression ast.Expr
		expr, ok := stmt.(*ast.ExprStmt)
		if !ok {
			assign, ok := stmt.(*ast.AssignStmt)
			if !ok || (assign != nil && len(assign.Rhs) != 1) {
				continue
			}
			expression = assign.Rhs[0]
		}

		/*
			the request body parsing can be done in many ways
			1) {
				decoder := json.NewDecoder(r.Body)
				decoder.Decode(req)

				and similarly for response
			}

			2) {
				json.NewDecoder(r.Body).Decode(req)
			}

			3) {
				body, err := io.ReadAll(r.Body)

				err := json.Unmarshal(body, req)

				for response writing,
				data, err := json.Marshal(res)
			}
		*/
		var call *ast.CallExpr

		if expression != nil {
			call, ok = expression.(*ast.CallExpr)
			if !ok {
				continue
			}
		} else {
			call, ok = expr.X.(*ast.CallExpr)
			if !ok {
				continue
			}
		}

		sl, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		id, ok := sl.X.(*ast.Ident)
		if !ok {
			call, ok := sl.X.(*ast.CallExpr)
			if ok {
				sl2, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				id, ok := sl2.X.(*ast.Ident)
				if !ok {
					continue
				}

				fmt.Println(sl.Sel.Name, sl2.Sel.Name, id.Name)
			}
			continue
		}

		fmt.Println(sl.Sel.Name, id.Name)
	}
}

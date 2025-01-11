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
		expr, ok := stmt.(*ast.ExprStmt)
		if !ok {
			continue
		}

		sl, ok := expr.X.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		id, ok := sl.X.(*ast.Ident)
		if !ok {
			continue
		}

		fmt.Println(sl.Sel.Name, id.Name)
	}
}

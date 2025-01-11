package parser

import (
	"fmt"
	"gen-doc/types"
	"gen-doc/utils"
	"strings"
)

// function used to extract all @ symbols
// used to represent a route
func (p *Parser) extractDocs() {
	for pkgname, fns := range p.handler_funcs {
		p.log.Println(pkgname)
		for _, fn := range fns {
			res := &types.ExtractRouteInfo{}

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

			fmt.Println(res)
		}
	}
}

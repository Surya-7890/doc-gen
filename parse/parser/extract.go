package parser

// function used to extract all @ symbols
// used to represent a route
func (p *Parser) extractDocs() {
	for pkgname, fns := range p.handler_funcs {
		p.log.Println(pkgname)
		for _, fn := range fns {
			for _, doc := range fn.Doc.List {
				p.log.Println(doc.Text)
			}
		}
	}
}

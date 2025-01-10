package parser

// function used to extract all @ symbols
// used to represent a route
func (p *Parser) extractDocs() {
	for _, fns := range p.handler_funcs {
		for _, fn := range fns {
			p.log.Println(fn.Doc.List)
		}
	}
}

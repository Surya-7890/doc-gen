package parser

// traverses from the main function through routers (*http.ServeMux)
// and through routers from other packages till a HandlerFunc is reached
func (p *Parser) traverseFunc() {
	for k, v := range p.fn_decls {
		for _, val := range v {
			p.log.Println(k, ":", val.Name.Name)
		}
	}
}

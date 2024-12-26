package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sync"
)

type Parser struct {
	log           *log.Logger
	call_exprs    []*ast.CallExpr
	fn_decls      map[string][]*ast.FuncDecl
	handler_funcs map[string][]*ast.FuncDecl
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{
		log:           logger,
		fn_decls:      make(map[string][]*ast.FuncDecl),
		handler_funcs: make(map[string][]*ast.FuncDecl),
	}
}

func (p *Parser) ParsePackages(files_map map[string][]string) {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	for k, v := range files_map {
		fileSet := token.NewFileSet()
		mode := parser.AllErrors | parser.ParseComments

		// maps [package_name] -> []FuncDecls (function declarations in the package)
		decl_arr := []*ast.FuncDecl{}
		p.fn_decls[k] = decl_arr

		for _, file := range v {
			wg.Add(1)
			go func(wg *sync.WaitGroup, mx *sync.Mutex) {
				defer wg.Done()
				fns, calls := p.parseSingleFile(file, fileSet, mode)

				mx.Lock()
				p.fn_decls[k] = append(p.fn_decls[k], fns...)
				p.call_exprs = append(p.call_exprs, calls...)
				mx.Unlock()

			}(wg, mx)
		}
	}
	wg.Wait()

	// HandleFuncs must be separated from the
	// other funcs and ancestral tree must be constructed
	p.filterHandlerFuncs()

	// todo ~ to generate the ancestral tree of handler_funcs

	p.traverseFunc()
}

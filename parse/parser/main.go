package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sync"
)

type Parser struct {
	log        *log.Logger
	call_exprs []*ast.CallExpr
	fn_decls   []*ast.FuncDecl
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{
		log: logger,
	}
}

func (p *Parser) ParsePackages(files_map map[string][]string) {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	for _, v := range files_map {
		fileSet := token.NewFileSet()
		mode := parser.AllErrors | parser.ParseComments

		for _, file := range v {
			wg.Add(1)
			go func(wg *sync.WaitGroup, mx *sync.Mutex) {
				defer wg.Done()
				fns, calls := p.parseSingleFile(file, fileSet, mode)

				mx.Lock()
				p.fn_decls = append(p.fn_decls, fns...)
				p.call_exprs = append(p.call_exprs, calls...)
				mx.Unlock()

			}(wg, mx)
		}
	}
	wg.Wait()

	// HandleFuncs must be separated from the
	// other funcs and ancestral tree must be constructed
	p.filterHandlerFuncs()
}

package parser

import (
	"fmt"
	"gen-doc/types"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sync"
)

type Parser struct {
	log        *log.Logger
	filenames  chan string
	call_exprs []*ast.CallExpr
	fn_decls   []*ast.FuncDecl
}

func NewParser(logger *log.Logger, channel chan string) *Parser {
	return &Parser{
		log:       logger,
		filenames: channel,
	}
}

func (p *Parser) WaitForFiles() {
	fileSet := token.NewFileSet()
	mode := parser.AllErrors | parser.ParseComments

	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	for filename := range p.filenames {
		// close the channel in case
		// types.CHANNEL_CLOSE is sent into
		// the channel from the scanner
		if filename == types.CHANNEL_CLOSE {
			close(p.filenames)
			break
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, mx *sync.Mutex) {
			defer wg.Done()
			fns, calls := p.parseSingleFile(filename, fileSet, mode)

			mx.Lock()
			p.fn_decls = append(p.fn_decls, fns...)
			p.call_exprs = append(p.call_exprs, calls...)
			mx.Unlock()

		}(wg, mx)
	}

	wg.Wait()

	// HandleFuncs must be separated from the
	// other funcs and ancestral tree must be constructed
	// for _, fn := range p.fn_decls {
	// 	fmt.Println(fn.Name.Name)
	// }

	//
	for _, call := range p.call_exprs {

		sl, ok := call.Fun.(*ast.SelectorExpr)
		if ok {
			id, ok := sl.X.(*ast.Ident)
			if ok {
				fmt.Println(id.Name + "::" + sl.Sel.Name)
			}
			continue
		}

		id, ok := call.Fun.(*ast.Ident)
		if ok {
			fmt.Println(id.Name)
		}
	}
}

package parser

import (
	"gen-doc/types"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sync"

	"golang.org/x/tools/go/packages"
)

type Parser struct {
	log           *log.Logger
	fn_decls      map[*types.MapKey][]*ast.FuncDecl
	handler_funcs map[*types.MapKey][]*ast.FuncDecl
	pkg_map       map[*types.MapKey]*packages.Package
	pkg_keys      []*types.MapKey
	dir_name      string
}

func NewParser(logger *log.Logger, dir_name string) *Parser {
	return &Parser{
		log:           logger,
		dir_name:      dir_name,
		fn_decls:      make(map[*types.MapKey][]*ast.FuncDecl),
		handler_funcs: make(map[*types.MapKey][]*ast.FuncDecl),
	}
}

func (p *Parser) ParsePackages(files_map map[*types.MapKey]*packages.Package) {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	mode := parser.AllErrors | parser.ParseComments

	p.pkg_map = files_map
	keys := []*types.MapKey{}
	for k := range p.pkg_map {
		keys = append(keys, k)
	}
	p.pkg_keys = keys

	for k, v := range files_map {

		fileSet := token.NewFileSet()

		// maps [package_name] -> []FuncDecls (function declarations in the package)
		p.fn_decls[k] = []*ast.FuncDecl{}

		for _, file := range v.GoFiles {
			ast_file, err := parser.ParseFile(fileSet, file, nil, mode)
			if err != nil {
				p.log.Fatal(err.Error())
			}

			wg.Add(1)
			go func(wg *sync.WaitGroup, mx *sync.Mutex) {
				defer wg.Done()
				fns := p.parseSingleFile(ast_file)

				mx.Lock()
				p.fn_decls[k] = append(p.fn_decls[k], fns...)
				mx.Unlock()

			}(wg, mx)
		}
	}
	wg.Wait()

	// HandleFuncs must be separated from the
	// other funcs and ancestral tree must be constructed
	p.filterHandlerFuncs()

	p.extractDocs()
}

package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

type IParser interface {
	Parse()
	ParseSingleFile(file_name string)
}

type Parser struct {
	Log *log.Logger
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{
		Log: logger,
	}
}

// gets directory name and returns a map
// the map contains filename as key and *ast.File as value
func (p *Parser) Parse() {

	// get the current working directory for parsing
	dir_name, err := os.Getwd()
	if err != nil {
		err_message := fmt.Sprintf("error while getting current working directory: %s", err.Error())
		p.Log.Fatal(err_message)
	}

	// returns a list of all files and folders in cwd
	// parse all files in the directory
	files, err := os.ReadDir(dir_name)
	if err != nil {
		err_message := fmt.Sprintf("error while parsing directory (%s): %s", dir_name, err.Error())
		p.Log.Fatal(err_message)
	}

	wg := &sync.WaitGroup{}
	for _, val := range files {

		// skips if a folder or any file other than a go file
		if val.IsDir() || filepath.Ext(val.Name()) != ".go" {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			p.ParseSingleFile(val.Name()) // pass wg as second param
		}(wg)
	}
	wg.Wait()
}

// gets filename as parameter
// *ast.File is then traversed
func (p *Parser) ParseSingleFile(file_name string) { // accept wg as second param
	file_set := token.NewFileSet()

	file, err := parser.ParseFile(file_set, file_name, nil, parser.AllErrors)
	if err != nil {
		err_message := fmt.Sprintf("error while parsing file(%s): %s", file_name, err.Error())
		p.Log.Fatal(err_message)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if !p.isRouteHandler(fn) {
			return true
		}

		return true
	})

	// wg.Done()
}

// checks if a function is a route handler
// reads the params to verify
// route handlers look like: HandlerName(res http.ResponseWriter, req *http.Request)
func (p *Parser) isRouteHandler(fn *ast.FuncDecl) bool {

	params := fn.Type.Params.List
	if len(params) != 2 {
		return false
	}

	res_type := reflect.TypeOf(params[0])
	req_type := reflect.TypeOf(params[1])

	if !res_type.Implements(reflect.TypeOf((http.ResponseWriter)(nil)).Elem()) {
		return false
	}

	if req_type.Kind() != reflect.TypeOf(&http.Request{}).Kind() {
		return false
	}

	return true
}

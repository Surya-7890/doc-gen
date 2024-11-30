package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type IParser interface {
	Parse()
	ParseSingleFile(file_name string)
}

type Parser struct {
	IParser
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
			messaage := fmt.Sprintf("function %s is not a route handler", fn.Name.Name)
			p.Log.Println(messaage)
			return true
		}

		message := fmt.Sprintf("%s is a route handler", fn.Name.Name)
		p.Log.Println(message)

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

	fmt.Println("1", params[0].Type)

	switch t := params[0].Type.(type) {
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok && x.Name == "http" && t.Sel.Name == "ResponseWriter" {
			p.Log.Println("First Parameter is http.ResponseWriter")
		}
	default:
		return false
	}

	switch t := params[1].Type.(type) {
	case *ast.StarExpr:
		val, ok := t.X.(*ast.SelectorExpr)
		if !ok {
			return false
		}

		if x, ok := val.X.(*ast.Ident); ok && x.Name == "http" && val.Sel.Name == "Request" {
			p.Log.Println("Second Parameter is *http.Request")
			return true
		}

		return false
	default:
		return false
	}
}

// searches route handler body for req.Body unmarshalling
// gets the type of req.Body
// generates json based on those parameters
func (p *Parser) generateSwaggerJSON() {}

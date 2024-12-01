package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type IParser interface {
	Parse() error
	ParseSingleFile(file_name string)
	IsRouteHandler(*ast.FuncDecl) bool
}

type Parser struct {
	IParser
	Log            *log.Logger
	ParsedFileChan chan []*ast.FuncDecl
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{
		Log:            logger,
		ParsedFileChan: make(chan []*ast.FuncDecl, 15),
	}
}

// gets directory name and returns a map
// the map contains filename as key and *ast.File as value
func (p *Parser) Parse() {

	files, err := p.ParseDir()
	if err != nil {
		p.Log.Println(err.Error())
		return
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

func (p *Parser) ParseDir() ([]fs.DirEntry, error) {
	// get the current working directory for parsing
	dir_name, err := os.Getwd()
	if err != nil {
		err_message := fmt.Sprintf("error while getting current working directory: %s", err.Error())
		p.Log.Println(err_message)
		return nil, err
	}

	// returns a list of all files and folders in cwd
	// parse all files in the directory
	files, err := os.ReadDir(dir_name)
	if err != nil {
		err_message := fmt.Sprintf("error while parsing directory (%s): %s", dir_name, err.Error())
		p.Log.Println(err_message)
		return nil, err
	}
	return files, nil
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

	arr := p.Traverse(file)

	p.ParsedFileChan <- arr
	time.Sleep(2 * time.Second)
	close(p.ParsedFileChan)

	// wg.Done()
}

// checks if a function is a route handler
// reads the params to verify
// route handlers look like: HandlerName(res http.ResponseWriter, req *http.Request)
func (p *Parser) IsRouteHandler(fn *ast.FuncDecl) bool {

	params := fn.Type.Params.List
	if len(params) != 2 {
		return false
	}

	switch t := params[0].Type.(type) {
	case *ast.SelectorExpr:
		x, ok := t.X.(*ast.Ident)
		if !ok || x.Name != "http" || t.Sel.Name != "ResponseWriter" {
			return false
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
			return true
		}

		return false
	default:
		return false
	}
}

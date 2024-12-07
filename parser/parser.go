package parser

import (
	"gen-doc/parser/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sync"
)

// with all .go files, traverses the AST
// stores all function declarations in p.FnMap
func (p *Parser) Parse() {
	dir, err := os.Getwd()

	scannerWg := &sync.WaitGroup{}
	Scanner := scanner.Scanner{
		Log:         p.log,
		FilesChan:   p.FilesChan,
		FuncDeclMap: make(map[string][]*ast.FuncDecl),
	}

	scannerWg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		Scanner.WaitForFiles()
	}(scannerWg)

	if err != nil {
		p.log.Fatal(err.Error())
	}

	files := p.getAllFiles(dir)

	wg := &sync.WaitGroup{}
	for _, file_name := range files {
		wg.Add(1)
		go func(file_name string) {
			defer wg.Done()
			p.parseSingleFile(file_name)
		}(file_name)
	}
	wg.Wait()
	close(p.FilesChan)
	scannerWg.Wait()
}

// parses single .go file to find all function declarations
func (p *Parser) parseSingleFile(file_name string) {
	file_set := token.NewFileSet()
	file, err := parser.ParseFile(file_set, file_name, nil, parser.ParseComments)
	if err != nil {
		p.log.Fatal(err.Error())
	}
	if file != nil {
		p.FilesChan <- file
	}
}

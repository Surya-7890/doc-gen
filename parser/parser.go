package parser

import (
	"go/parser"
	"go/token"
	"os"
	"sync"
)

// with all .go files, traverses the AST
// stores all function declarations in p.FnMap
func (p *Parser) Parse() {
	dir, err := os.Getwd()
	if err != nil {
		p.log.Fatal(err.Error())
	}

	files := p.getAllFiles(dir)

	wg := &sync.WaitGroup{}
	for _, file_name := range files {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			p.parseSingleFile(file_name, wg)
		}(wg)
	}
	wg.Wait()
}

// parses single .go file to find all function declarations
func (p *Parser) parseSingleFile(file_name string, wg *sync.WaitGroup) {
	file_set := token.NewFileSet()
	file, err := parser.ParseFile(file_set, file_name, nil, parser.ParseComments)
	if err != nil {
		p.log.Fatal(err.Error())
	}
	if file != nil {
		p.FilesChan <- file
	}
	wg.Done()
}

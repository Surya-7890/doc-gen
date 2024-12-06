package parser

import (
	"fmt"
	"os"
)

// with all .go files, traverses the AST
// stores all function declarations in p.FnMap
func (p *Parser) Parse() {
	dir, err := os.Getwd()
	if err != nil {
		p.log.Fatal(err.Error())
	}

	files := p.getAllFiles(dir)

	for _, val := range files {
		fmt.Println(val)
	}
}

// parses single .go file to find all function declarations
func (p *Parser) parseSingleFile() {}

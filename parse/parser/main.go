package parser

import (
	"go/parser"
	"go/token"
	"log"
)

type Parser struct {
	log       *log.Logger
	filenames chan string
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
	for filename := range p.filenames {
		p.parseSingleFile(filename, fileSet, mode)
	}
}

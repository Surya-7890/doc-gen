package parser

import (
	"gen-doc/parser/scanner"
	"go/ast"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

type IParser interface {
	getAllFiles() []fs.DirEntry
	Parse()
	parseSingleFile()
}

type Parser struct {
	IParser
	log       *log.Logger
	FilesChan chan *ast.File
}

func NewParser(logger *log.Logger) *Parser {
	channel := make(chan *ast.File, 10)
	Scanner := scanner.Scanner{
		Log:       logger,
		FilesChan: channel,
	}
	go Scanner.WaitForFiles()
	return &Parser{
		log:       logger,
		FilesChan: channel,
	}
}

// gets all .go files in the current working directory
func (p *Parser) getAllFiles(dir_name string) []string {
	files, err := os.ReadDir(dir_name)
	if err != nil {
		p.log.Fatal(err.Error())
	}

	var arr []string

	for _, entry := range files {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			arr = append(arr, p.getAllFiles(path.Join(dir_name, entry.Name()))...)
		}
		if path.Ext(entry.Name()) != ".go" {
			continue
		}
		arr = append(arr, path.Join(dir_name, entry.Name()))
	}

	return arr
}
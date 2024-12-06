package parser

import (
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
	findHandlers()
	findHandlerFuncs()
}

type Parser struct {
	IParser
	log   *log.Logger
	fnMap map[string]ast.Node
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{
		log:   logger,
		fnMap: make(map[string]ast.Node),
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

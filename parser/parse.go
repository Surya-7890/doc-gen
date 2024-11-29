package parser

import (
	"fmt"
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
			p.parseSingleFile(val.Name(), wg)
		}(wg)
	}
	wg.Wait()

}

// gets filename as parameter
// *ast.File is then traversed
func (p *Parser) parseSingleFile(file_name string, wg *sync.WaitGroup) {

}

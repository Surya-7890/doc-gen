package test

import (
	"fmt"
	"gen-doc/parser"
	"log"
	"os"
	"testing"
)

func TestParseSingleFile(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	Parser.ParseSingleFile("./functions.go", nil)

	for val := range Parser.ParsedFileChan {
		if val == nil || len(val) == 0 {
			t.Fail()
		} else {
			for _, v := range val {
				if v != nil {
					fmt.Println(v.Name.Name)
					close(Parser.ParsedFileChan)
				}
			}
		}
	}
}

func TestParseDir(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	// get the current working directory for parsing
	dir_name, err := os.Getwd()
	if err != nil {
		err_message := fmt.Sprintf("error while getting current working directory: %s", err.Error())
		Parser.Log.Println(err_message)
	}

	files, err := Parser.ParseDir(dir_name)
	if err != nil || len(files) == 0 {
		t.FailNow()
	}
}

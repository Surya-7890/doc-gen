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

	Parser.ParseSingleFile("./functions.go")

	for val := range Parser.ParsedFileChan {
		if val == nil || len(val) == 0 {
			t.Fail()
		} else {
			for _, v := range val {
				if v != nil {
					fmt.Println(v.Name.Name)
				}
			}
		}
	}
}

func TestParseDir(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	files, err := Parser.ParseDir()
	if err != nil || len(files) == 0 {
		t.FailNow()
	}
}

package main

import (
	"gen-doc/parser"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	Parser.Parse()
}

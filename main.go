package main

import (
	"gen-doc/parser"
	"log"
)

func main() {
	logger := log.Default()

	Parser := parser.NewParser(logger)

	Parser.Parse()
}

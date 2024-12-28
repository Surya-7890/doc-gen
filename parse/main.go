package parse

import (
	"gen-doc/parse/parser"
	"gen-doc/parse/scanner"
	"log"
)

func Parse() {
	logger := log.Default()

	Parser := parser.NewParser(logger)
	Scanner := scanner.NewScanner(logger)

	files_map := Scanner.GetAllFiles()
	Parser.ParsePackages(files_map)
}

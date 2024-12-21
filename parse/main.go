package parse

import (
	"gen-doc/parse/parser"
	"gen-doc/parse/scanner"
	"log"
)

func Parse() {
	logger := log.Default()
	channel := make(chan string, 10)
	Parser := parser.NewParser(logger, channel)
	Scanner := scanner.NewScanner(logger, channel)

	go Parser.WaitForFiles()

	Scanner.GetAllFiles()
}

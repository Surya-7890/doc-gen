package parse

import (
	"gen-doc/parse/parser"
	"gen-doc/parse/scanner"
	"log"
	"sync"
)

func Parse() {
	logger := log.Default()
	channel := make(chan string, 10)
	Parser := parser.NewParser(logger, channel)
	Scanner := scanner.NewScanner(logger, channel)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		Parser.WaitForFiles()
	}(wg)

	Scanner.GetAllFiles()

	wg.Wait()
}

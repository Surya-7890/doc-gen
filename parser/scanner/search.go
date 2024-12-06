package scanner

import (
	"go/ast"
	"log"
	"sync"
)

type IScanner interface {
	WaitForFiles()
	traverse()
	findHandlers()
	findHandlerFuncs()
}

type Scanner struct {
	IScanner
	Log       *log.Logger
	FilesChan chan *ast.File
}

func (s *Scanner) WaitForFiles() {
	wg := &sync.WaitGroup{}
	for file := range s.FilesChan {
		if file.Name.Name == "main" {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				s.traverse(file, wg)
			}(wg)
		}
	}
	wg.Wait()
}

func (s *Scanner) traverse(file *ast.File, wg *sync.WaitGroup) {
	ast.Inspect(file, func(n ast.Node) bool {
		s.findHandlers(n)
		return true
	})
	wg.Done()
}

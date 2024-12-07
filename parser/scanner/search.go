package scanner

import (
	"fmt"
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
	i := 0
	for file := range s.FilesChan {
		i++
		if file.Name.Name == "main" {
			wg.Add(1)
			go func(file *ast.File) {
				defer wg.Done()
				s.traverse(file)
			}(file)
		}
	}
	wg.Wait()
	fmt.Println(i)
}

func (s *Scanner) traverse(file *ast.File) {
	ast.Inspect(file, func(n ast.Node) bool {
		s.findHandlers(n)
		return true
	})
}

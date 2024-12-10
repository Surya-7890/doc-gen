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
	Log         *log.Logger
	FilesChan   chan *ast.File
	FuncDeclMap map[string][]*ast.FuncDecl
}

func (s *Scanner) WaitForFiles() {
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	for file := range s.FilesChan {
		wg.Add(1)
		go func(file *ast.File) {
			defer wg.Done()
			s.getAllFuncDecls(file, mx)
		}(file)
	}
	wg.Wait()

	for _, fn := range s.FuncDeclMap["main"] {
		wg.Add(1)
		go func(fn *ast.FuncDecl) {
			defer wg.Done()
			args := s.findHandlers(fn)
			if len(args) > 0 {
				slct_stmt, ok := args[1].(*ast.SelectorExpr)
				if ok {
					fmt.Println("sel:", slct_stmt.X, " name:", slct_stmt.Sel.Name)
					return
				}
				fmt.Println(args[1])
			}
		}(fn)
	}
	wg.Wait()
}

// func (s *Scanner) traverse(fn *ast.FuncDecl, mx *sync.Mutex) {
// 	ast.Inspect(fn, func(n ast.Node) bool {
// 		return true
// 	})
// }

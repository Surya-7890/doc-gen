package scanner

import (
	"go/ast"
	"sync"
)

func (s *Scanner) getAllFuncDecls(file *ast.File, mx *sync.Mutex) {
	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		mx.Lock()
		s.FuncDeclMap[file.Name.Name] = append(s.FuncDeclMap[file.Name.Name], fn)
		mx.Unlock()

		return true
	})
}

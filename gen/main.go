package gen

import (
	"fmt"
	"go/ast"
)

type Gen struct {
	RoutesChan chan []*ast.FuncDecl
}

func NewGenerator(channel chan []*ast.FuncDecl) *Gen {
	return &Gen{
		RoutesChan: channel,
	}
}

func (g *Gen) WaitForFiles() {
	for val := range g.RoutesChan {
		if len(val) == 0 {
			continue
		}
		fmt.Println(val)
		for _, v := range val {
			g.constructRouteInfo(v)
		}
	}
}

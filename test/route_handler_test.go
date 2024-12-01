package test

import (
	"gen-doc/parser"
	"log"
	"os"
	"testing"
)

func TestRouteHandler(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	fn, ok := Map["RouteHandler"]
	if !ok {
		t.FailNow()
	}

	ok = Parser.IsRouteHandler(fn)
	if !ok {
		t.FailNow()
	}

}

func TestIncorrectParams(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	fn, ok := Map["IncorrectParams"]
	if !ok {
		t.FailNow()
	}

	ok = Parser.IsRouteHandler(fn)
	if ok {
		t.FailNow()
	}

}

func TestIncorrectFirstParam(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	fn, ok := Map["IncorrectFirstParam"]
	if !ok {
		t.FailNow()
	}

	ok = Parser.IsRouteHandler(fn)
	if ok {
		t.FailNow()
	}
}

func TestIncorrectSecondParam(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	fn, ok := Map["IncorrectSecondParam"]
	if !ok {
		t.FailNow()
	}

	ok = Parser.IsRouteHandler(fn)
	if ok {
		t.FailNow()
	}
}

func TestIncorrectNoOfParams(t *testing.T) {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := parser.NewParser(logger)

	fn, ok := Map["IncorrectNoOfParams"]
	if !ok {
		t.FailNow()
	}

	ok = Parser.IsRouteHandler(fn)
	if ok {
		t.FailNow()
	}
}

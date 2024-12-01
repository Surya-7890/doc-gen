package main

import (
	"gen-doc/app"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[gen-doc]: ", log.Flags())
	Parser := app.NewParser(logger)

	Parser.Parse()
}

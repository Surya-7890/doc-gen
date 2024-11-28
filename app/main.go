package app

import (
	"gen-doc/cli"
)

type IApplication interface {
	Init() *Application
}

type Application struct {
	Cli *cli.CLI
}

func Init() *Application {
	return &Application{}
}

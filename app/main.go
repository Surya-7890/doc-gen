package app

type IApplication interface {
	Init() *Application
}

type Application struct{}

func Init() *Application {
	return &Application{}
}

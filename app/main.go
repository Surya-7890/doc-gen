package app

type IApplication interface {
	Init() *Application
}

type Application struct {
	IApplication
}

func Init() *Application {
	return &Application{}
}

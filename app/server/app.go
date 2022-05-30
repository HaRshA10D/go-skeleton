package server

import (
	"github.com/user-name/skeleton-name/handler"

	"github.com/user-name/skeleton-name/config"
	pkghttp "github.com/user-name/skeleton-name/pkg/http"
	"github.com/user-name/skeleton-name/service"
	"github.com/user-name/skeleton-name/store"
)

type App struct {
	config     *config.AppConfig
	server     pkghttp.Server
	repository store.Repository
	useCases   service.UseCases
}

func InitializeApp() (*App, error) {
	app := App{}
	app.config = config.Current

	app.repository = app.setupRepository()
	app.useCases = app.setupServices()

	router := handler.NewRouter()

	app.server = pkghttp.NewServer(app.config.Server, router)
	return &app, nil
}

func (a *App) Start() {
	a.server.Start()
}

func (a *App) setupServices() service.UseCases {
	return service.UseCases{}
}

func (a *App) setupRepository() store.Repository {
	return store.Repository{}
}

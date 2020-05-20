package app

import (
    "url-shortener/api"
    "url-shortener/interfaces"
    "url-shortener/storage"
)

type App struct {
    config *interfaces.Config
    api    *api.Api
    store  storage.Store
}

func NewApp(config *interfaces.Config) *App {
    return &App{
        config: config,
    }
}

func (a *App) Config() *interfaces.Config {
    return a.config
}

func (a *App) SetApi(apiApp *api.Api) {
    a.api = apiApp
}

func (a *App) SetStore(store storage.Store) {
    a.store = store
}

func (a *App) Store() storage.Store {
    return a.store
}

func (a *App) Run() error {
    return a.api.Run()
}

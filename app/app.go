package app

import (
    "url-shortener/api"
    "url-shortener/interfaces"
)

type App struct {
    config *interfaces.Config
    api    *api.Api
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

func (a *App) Run() error {
    return a.api.Run()
}

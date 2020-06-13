package interfaces

import (
    "url-shortener/models"
    "url-shortener/storage"
)

type App interface {
    Config() *Config
    Run() error
    Store() storage.Store

    SaveUrl(*models.Url) (*models.Url, error)
    FindUrl(string) (*models.Url, error)
}

type Config struct {
    Host   string
    Port   int
    DbHost string
    DBPort int
    DbName string
}

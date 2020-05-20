package interfaces

import "url-shortener/storage"

type App interface {
    Config() *Config
    Run() error
    Store() storage.Store
}

type Config struct {
    Host   string
    Port   int
    DbHost string
    DBPort int
    DbName string
}

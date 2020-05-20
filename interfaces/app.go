package interfaces

type App interface {
    Config() *Config
    Run() error
}

type Config struct {
    Host   string
    Port   int
    DbHost string
    DBPort int
    DbName string
}

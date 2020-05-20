package commands

import (
    "github.com/sirupsen/logrus"
    "url-shortener/api"
    "url-shortener/app"
    "url-shortener/interfaces"
)

type ServerCommand struct {
    Host string        `long:"host" env:"HOST" default:"127.0.0.1" description:"server host"`
    Port int           `long:"port" env:"PORT" default:"6969" description:"server port"`
    Db   DatabaseGroup `group:"db" namespace:"db" env-namespace:"DB"`
}

type DatabaseGroup struct {
    Name string `long:"name" env:"NAME" default:"url-shortener" description:"DB name"`
    Host string `long:"host" env:"HOST" default:"127.0.0.1" description:"DB host"`
    Port int    `long:"port" env:"PORT" default:"27017" description:"DB port"`
}


func (s *ServerCommand) Execute(args []string) error {
    logrus.Info("Running server command")

    app := s.initApp()

    app.Run()

    return nil
}

func (s *ServerCommand) initApp() interfaces.App {
    serverApp := app.NewApp(&interfaces.Config{
        Host: s.Host,
        Port: s.Port,
        DbName: s.Db.Name,
        DbHost: s.Db.Host,
        DBPort: s.Db.Port,
    })

    apiApp := api.NewApi(serverApp)
    apiApp.InitApi()

    serverApp.SetApi(apiApp)

    return serverApp
}

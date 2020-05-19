package commands

import (
    "github.com/sirupsen/logrus"
)

type ServerCommand struct {
    Host string        `long:"host" env:"HOST" default:"127.0.0.1" description:"server host"`
    Port int           `long:"port" env:"PORT" default:"6969" description:"server port"`
    Db   DatabaseGroup `group:"db" namespace:"db" env-namespace:"DB"`
}

type DatabaseGroup struct {
    Name string `long:"name" env:"NAME" default:"sea-battle" description:"DB name"`
    Host string `long:"host" env:"HOST" default:"127.0.0.1" description:"DB host"`
    Port int    `long:"port" env:"PORT" default:"27017" description:"DB port"`
}


func (s *ServerCommand) Execute(args []string) error {
    logrus.Info("Running server command")
    
    return nil
}

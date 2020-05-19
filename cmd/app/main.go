package main

import (
    "github.com/jessevdk/go-flags"
    "github.com/sirupsen/logrus"
    "os"
    "url-shortener/cmd/commands"
)

type Options struct {
    ServerCmd  commands.ServerCommand  `command:"server"`
}

func main() {
    configureLogger()
    logrus.Info("Starting backend app...")

    var opts Options
    p := flags.NewParser(&opts, flags.Default)
    if _, err := p.Parse(); err != nil {
        if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
            os.Exit(0)
        } else {
            os.Exit(1)
        }
    }
}

func configureLogger() {
    logrus.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
}

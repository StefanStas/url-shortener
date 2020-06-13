package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "url-shortener/interfaces"
)

type Api struct {
    app    interfaces.App
    router *gin.Engine
}

func NewApi(app interfaces.App) *Api {
    return &Api{
        app: app,
        router: gin.Default(),
    }
}

func (a *Api) InitApi() {
    a.createRoutes()
}

func (a *Api) createRoutes() {
    a.router.GET("/ping", a.PingCtrl)
    api := a.router.Group("/api")
    {
        api.GET("/urls/:hash", a.RedirectUrlCtrl)
        api.POST("/urls", a.SaveUrlValidator, a.SaveUrlCtrl)
    }
}

func (a *Api) Run() error {
    conf := a.app.Config()
    return a.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}

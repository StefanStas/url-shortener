package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (a *Api) PingCtrl(c *gin.Context) {
    urls, err := a.app.Store().UrlStore().GetUrls()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err,
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
        "urls": urls,
    })
}

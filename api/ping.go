package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (a *Api) PingCtrl(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

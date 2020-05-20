package api

import "github.com/gin-gonic/gin"

func (a *Api) PingCtrl(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}

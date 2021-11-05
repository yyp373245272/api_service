package handler

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
    log.Println(c.Request.Header)
    c.String(http.StatusOK, "pong")
}
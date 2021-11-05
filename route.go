package main

import (
    "api_service/handler"

    "github.com/gin-gonic/gin"
)

func register(s *gin.Engine){
    s.GET("/ping",handler.Ping)
}
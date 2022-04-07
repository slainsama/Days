package main

import (
	"Days/controller"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("login", controller.Login)
	return router
}

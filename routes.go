package main

import (
	"Days/controller"
	"Days/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("login", controller.Login)
	router.GET("index", middleware.IsAuthed, controller.IndexGetImage)
	return router
}

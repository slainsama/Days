package controller

import (
	"Days/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func Login(context *gin.Context) {
	username := config.AdminConfig.GetString("admin.username")
	password := config.AdminConfig.GetString("admin.password")
	var PostData LoginInfo
	err := context.ShouldBindJSON(&PostData)
	if err != nil {
		fmt.Println("请求错误")
	}
	if PostData.UserName == username && PostData.PassWord == password {
		return
	}
}

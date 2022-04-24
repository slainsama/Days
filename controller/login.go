package controller

import (
	"Days/global"
	"Days/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func Login(context *gin.Context) {
	username := global.Config.GetString("admin.username")
	password := global.Config.GetString("admin.password")
	fmt.Println(username)
	fmt.Println(password)
	var PostData LoginInfo
	err := context.ShouldBindJSON(&PostData)
	if err != nil {
		fmt.Println("请求错误", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "wrong request"})
		return
	}
	fmt.Println(PostData)
	if PostData.UserName == username && PostData.PassWord == password {
		session, err := middleware.Store.Get(context.Request, "user")
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"error": "wrong"})
			return
		}
		session.Values["username"] = "admin"
		session.Save(context.Request, context.Writer)
		context.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "failed"})
		return
	}
}

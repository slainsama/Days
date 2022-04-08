package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("hereissecret"))

func IsAuthed(context *gin.Context) {
	session, _ := Store.Get(context.Request, "user")
	if session.Values["username"] != "admin" {
		context.JSON(http.StatusOK, gin.H{"msg": "err,No author"})
	} else {
		context.Next()
	}
}

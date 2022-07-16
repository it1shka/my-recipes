package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func errorMessageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		defer func() {
			session.Delete("error")
			session.Save()
		}()

		myerr := session.Get("error")
		var message string
		if myerr != nil {
			message = myerr.(string)
		} else {
			message = ""
		}
		ctx.Set("error", message)
		ctx.Next()
	}
}

func Setup(server *gin.Engine) {
	server.GET("/", indexHandler)

	{
		server.GET("/login", errorMessageMiddleware(), getLoginHandler)
		server.GET("/register", errorMessageMiddleware(), getRegisterHandler)
		server.POST("/login", postLoginHandler)
		server.POST("/register", postRegisterHandler)
		server.GET("/logout", getLogoutHander)
	}
}

func indexHandler(ctx *gin.Context) {
	// here goes the rest
	session := sessions.Default(ctx)
	username := session.Get("username")
	useremail := session.Get("useremail")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"username":  username,
		"useremail": useremail,
	})
}

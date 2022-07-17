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

func authRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userid := session.Get("userid")
		if userid == nil {
			session.Set("error", "Authentication required")
			session.Save()
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

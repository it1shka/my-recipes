package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

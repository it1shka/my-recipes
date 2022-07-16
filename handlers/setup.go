package handlers

import (
	"github.com/gin-gonic/gin"
)

func Setup(server *gin.Engine) {
	server.GET("/", indexHandler)

	{
		server.GET("/login", errorMessageMiddleware(), getLoginHandler)
		server.GET("/register", errorMessageMiddleware(), getRegisterHandler)
		server.GET("/logout", getLogoutHander)

		server.POST("/login", postLoginHandler)
		server.POST("/register", postRegisterHandler)
	}

	recipeRouter := server.Group("/recipe")
	{
		recipeRouter.GET("/add", authRequiredMiddleware(), getRecipeAddHandler)
	}
}
package handlers

import (
	"net/http"

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
		protected := recipeRouter.Group("/", authRequiredMiddleware())
		protected.GET("/add", errorMessageMiddleware(), getRecipeAddHandler)
		protected.POST("/add", postRecipeAddHandler)

		recipeRouter.GET("/:slug", getRecipeBySlugHandler)
	}

	server.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})
}

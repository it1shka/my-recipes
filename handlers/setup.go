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
		// protected routes
		protected := recipeRouter.Group("/", authRequiredMiddleware())
		protected.GET("/add", errorMessageMiddleware(), getRecipeAddHandler)
		protected.POST("/add", postRecipeAddHandler)

		// protected delete and edit
		managing := protected.Group("/:slug", ensureRecipeExistanceMiddleware(), ensureAuthorMiddleware())
		managing.GET("/delete", getRecipeDeleteHandler)
		managing.GET("/edit", errorMessageMiddleware(), getRecipeEditHandler)
		managing.POST("/edit", postRecipeEditHandler)

		// public routes
		recipeRouter.GET("/:slug", ensureRecipeExistanceMiddleware(), getRecipeBySlugHandler)
	}

	server.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})
}

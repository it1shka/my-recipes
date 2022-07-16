package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRecipeAddHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "recipe_form.html", gin.H{})
}

func postRecipeAddHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, world!")
}

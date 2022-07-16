package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRecipeAddHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "recipe_form.html", gin.H{})
}

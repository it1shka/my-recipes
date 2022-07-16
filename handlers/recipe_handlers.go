package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRecipeAddHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "You are allowed to visit the page!")
}

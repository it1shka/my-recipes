package myutils

import (
	"github.com/gin-gonic/gin"
	"it1shka.com/my-recipes/database"
)

func ExcludeRecipe(ctx *gin.Context) database.Recipe {
	maybeRecipe, _ := ctx.Get("recipe")
	recipe := maybeRecipe.(database.Recipe)
	return recipe
}

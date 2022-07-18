package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"it1shka.com/my-recipes/database"
	"it1shka.com/my-recipes/myutils"
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
			session.Set("error", "Auth error: authentication required")
			session.Save()
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func ensureRecipeExistanceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slug := ctx.Param("slug")
		recipe, exists := database.FindRecipeBySlug(slug)
		if !exists {
			ctx.HTML(http.StatusNotFound, "recipe_not_found.html", nil)
			ctx.Abort()
			return
		}

		ctx.Set("recipe", recipe)
		ctx.Next()
	}
}

func ensureAuthorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recipe := myutils.ExcludeRecipe(ctx)
		userid := sessions.Default(ctx).Get("userid")

		if userid != recipe.AuthorID {
			ctx.HTML(http.StatusForbidden, "403.html", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

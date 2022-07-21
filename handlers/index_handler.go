package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"it1shka.com/my-recipes/database"
	"it1shka.com/my-recipes/myutils"
)

func indexHandler(ctx *gin.Context) {
	maybePage := ctx.Query("page")
	page, err := strconv.Atoi(maybePage)
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	search := ctx.Query("search")
	recipes := database.FindRecipesByPage(page, search)
	pagination := myutils.GeneratePages(page)

	session := sessions.Default(ctx)
	username := session.Get("username")
	useremail := session.Get("useremail")

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"username":   username,
		"useremail":  useremail,
		"recipes":    myutils.PrepareRecipes(recipes),
		"page":       page,
		"pagination": pagination,
	})
}

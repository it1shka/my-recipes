package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gosimple/slug"
	"it1shka.com/my-recipes/database"
)

const RECIPE_SHORT_MESSAGE = `
	recipe title or description
	is too short:
	title should be at least 5 characters long;
	description should be at least 250 characters long
`

func getRecipeAddHandler(ctx *gin.Context) {
	errmsg := ctx.GetString("error")
	ctx.HTML(http.StatusOK, "recipe_form.html", gin.H{
		"error": errmsg,
	})
}

type RecipeFormData struct {
	Title       string `form:"title" bind:"required"`
	Description string `form:"description" bind:"required"`
}

func postRecipeAddHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	defer session.Save()

	recipeError := func(message string) {
		errmsg := fmt.Sprintf("Failed to add recipe: %s", message)
		session.Set("error", errmsg)
		ctx.Redirect(http.StatusFound, "/recipe/add")
	}

	var recipeData RecipeFormData
	if err := ctx.ShouldBind(&recipeData); err != nil {
		recipeError("form validation failed")
		return
	}

	if len(recipeData.Title) < 5 || len(recipeData.Description) < 250 {
		recipeError(RECIPE_SHORT_MESSAGE)
		return
	}

	title := recipeData.Title
	slug := slug.Make(title)

	if _, exists := database.FindRecipeBySlug(slug); exists {
		recipeError("recipe with the same or similar name already exists")
		return
	}

	description := recipeData.Description
	authorID := session.Get("userid").(uint)

	if _, err := database.CreateRecipe(slug, title, description, authorID); err != nil {
		recipeError("failed to add your recipe")
		return
	}

	redirectRoute := fmt.Sprintf("/recipe/%s", slug)
	ctx.Redirect(http.StatusFound, redirectRoute)
}

func getRecipeBySlugHandler(ctx *gin.Context) {

	slug := ctx.Param("slug")
	recipe, exists := database.FindRecipeBySlug(slug)
	if !exists {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	authorName := database.AuthorNameById(recipe.AuthorID)
	createdAt := recipe.CreatedAt.Format("2006-02-01")
	description := string(markdown.ToHTML([]byte(recipe.Description), nil, nil))
	currentUserId := sessions.Default(ctx).Get("userid").(uint)

	ctx.HTML(http.StatusOK, "recipe_page.html", gin.H{
		"title":       recipe.Title,
		"authorname":  authorName,
		"createdat":   createdAt,
		"description": template.HTML(description),

		"isAuthor": currentUserId == recipe.AuthorID,
	})
}

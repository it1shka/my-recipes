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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const PAGINATION_SIZE = 5

func GeneratePages(page int) []int {
	pages := make([]int, PAGINATION_SIZE)
	index := 0
	start := max(1, page-2)
	for i := start; i < start+PAGINATION_SIZE; i++ {
		pages[index] = i
		index++
	}
	return pages
}

type RecipeInfo struct {
	Slug, Title, AuthorName, CreatedAt string
}

func GetRecipeInfo(recipe database.Recipe) RecipeInfo {
	authorName := database.AuthorNameById(recipe.AuthorID)
	createdAt := recipe.CreatedAt.Format("2006-02-01")
	info := RecipeInfo{
		Slug:       recipe.Slug,
		Title:      recipe.Title,
		AuthorName: authorName,
		CreatedAt:  createdAt,
	}
	return info
}

func PrepareRecipes(recipes []database.Recipe) []RecipeInfo {
	infos := make([]RecipeInfo, len(recipes))
	for idx, val := range recipes {
		infos[idx] = GetRecipeInfo(val)
	}
	return infos
}

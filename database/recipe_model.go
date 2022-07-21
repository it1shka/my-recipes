package database

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Slug        string `gorm:"primaryIndex"`
	Title       string
	Description string
	AuthorID    uint
}

func CreateRecipe(slug, title, description string, authorID uint) (Recipe, error) {
	recipe := Recipe{
		Slug:        slug,
		Title:       title,
		Description: description,
		AuthorID:    authorID,
	}
	err := DB.Create(&recipe).Error
	return recipe, err
}

func FindRecipeBySlug(slug string) (recipe Recipe, exists bool) {
	err := DB.Where("slug = ?", slug).First(&recipe).Error
	exists = (err != gorm.ErrRecordNotFound)
	return
}

const PAGE_SIZE = 10

// func FindRecipesByPage(page int) []Recipe {
// 	offset := (page - 1) * PAGE_SIZE
// 	var recipes []Recipe
// 	DB.Offset(offset).Limit(PAGE_SIZE).Order("created_at desc").Find(&recipes)
// 	return recipes
// }

func FindRecipesByPage(page int, search string) []Recipe {
	offset := (page - 1) * PAGE_SIZE
	var recipes []Recipe

	query := DB // initialize a query
	query = query.Where("lower(title) like ?", "%"+search+"%")
	query = query.Offset(offset)
	query = query.Limit(PAGE_SIZE)
	query = query.Order("created_at desc")
	query.Find(&recipes)

	return recipes
}

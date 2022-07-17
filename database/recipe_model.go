package database

import "gorm.io/gorm"

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

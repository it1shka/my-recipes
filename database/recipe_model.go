package database

type Recipe struct {
	Slug        string `gorm:"primaryIndex"`
	Title       string
	Description string
	AuthorID    uint
}

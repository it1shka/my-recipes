package database

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey,autoIncrement"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func CreateUser(name, email, password string) (User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := DB.Create(&user).Error
	return user, err
}

func FindUserByEmail(email string) (user User, exists bool) {
	err := DB.Where("email = ?", email).First(&user).Error
	exists = (err != gorm.ErrRecordNotFound)
	return
}

func FindUserById(id uint) (user User, exists bool) {
	err := DB.Where("id = ?", id).First(&user).Error
	exists = (err != gorm.ErrRecordNotFound)
	return
}

func AuthorNameById(id uint) string {
	author, exists := FindUserById(id)
	if exists {
		return author.Name
	}
	return "Unknown"
}

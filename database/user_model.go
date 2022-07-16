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

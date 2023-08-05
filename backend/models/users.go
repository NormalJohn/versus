package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Users struct {
	gorm.Model
	Username string `gorm:"unique"` // This tag ensures usernames are unique in the table
	Nickname string `gorm:"unique"` // This tag ensures nicknames are unique in the table
	Password []byte // This should hold the hashed password, not the plaintext one
	Email    string `gorm:"type:varchar(100);unique"` // Setting type and ensuring email uniqueness
}

// CreateUser This function is used to create a new user in the database
func CreateUser(username string, nickname string, hashedPassword []byte, email string) bool {
	user := Users{Username: username, Nickname: nickname, Password: hashedPassword, Email: email}
	result := db.Create(&user)
	if result.Error != nil {
		return false
	}
	return true
}

func CompareHashAndContent(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// CheckLogin
func CheckLogin(username string, password string) (bool, error) {
	var user Users
	db.Where("username = ?", username).First(&user)
	if user.ID > 0 {
		if CompareHashAndContent(string(user.Password), password) {
			return true, nil
		} else {
			return false, errors.New("wrong")
		}
	}
	return false, errors.New("wrong")
}

// GetUserInfoByUsername
func GetUserInfoByUsername(username string) *Users {
	var user Users
	db.Select("id", "username", "nickname", "email").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return &user
	}
	return nil
}

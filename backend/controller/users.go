package controller

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// Register is the function that handles the POST request to /register
// It takes the username, nickname, password and email from the request body
// It then calls the CreateUser function from models/users.go
// backend always returns 200 HTTP codes, and it has a self-defined code in the response body
// If the user is created successfully, it returns 200  code and an ok.
// If the user is not created successfully, it returns a 400  code and a message contains specific error message.
func Register(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "Invalid form",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "Bad Request",
		})
		return
	}
	if models.CreateUser(form.Username, form.Nickname, hashedPassword, form.Email) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "User already exists",
		})
		return
	}

}

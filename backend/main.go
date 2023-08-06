package main

import (
	"backend/controller"
	"backend/middleware"
	"backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	jwt := middleware.JWTAuthMiddleWare("versus_123990hh_secret")
	router.POST("/login", jwt.LoginHandler)
	router.POST("/register", controller.Register)
	auth := router.Group("/auth")
	auth.Use(jwt.MiddlewareFunc())
	models.Connect()
	router.Run(":8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username != "admin" || password != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credential",
			})
			return
		}

		// FIXED DOMAIN
		c.SetCookie("session", "data", 3600, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		user, err := c.Cookie("session")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Cookie not found. Please login.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Dashboard accessed successfully",
			"user":    user,
		})
	})

	r.Run(":8080")
}

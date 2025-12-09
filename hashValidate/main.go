package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `form:"username" json:"username" binding:"required,min=3"`
	Password string `form:"password" json:"password" binding:"required,min=4"`
}

// Hardcoded hashed password for "1234"
var storedHashedPassword string

func main() {
	// Generate hash only once
	hash, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	storedHashedPassword = string(hash)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret123"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/login", LoginHandler)
	router.GET("/home", AuthMiddleware(), HomeHandler)

	router.Run(":8080")
}

func LoginHandler(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Username validation (hardcoded)
	if input.Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// Compare password with stored bcrypt hash
	if err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Save session
	session := sessions.Default(c)
	session.Set("user", input.Username)
	session.Save()

	c.JSON(200, gin.H{"message": "Login success"})
}

func HomeHandler(c *gin.Context) {
	user := sessions.Default(c).Get("user")
	c.JSON(200, gin.H{"message": "Welcome", "user": user})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user") == nil {
			c.JSON(401, gin.H{"error": "Login required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

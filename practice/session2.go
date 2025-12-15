package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    
    store := cookie.NewStore([]byte("secret123"))
    r.Use(sessions.Sessions("mysession",store))
    
    r.POST("/login", func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username == "admin" && password == "1234"{
            session := sessions.Default(c)
            session.Set("user", username)
            session.Save()
            
            c.JSON(http.StatusOK, gin.H{
                "message":"login succesfulll",
            })
            return
        }
        
        c.JSON(http.StatusUnauthorized, gin.H{
            "message":"invalid credential",
        })
    })
    
    r.GET("/login", func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        
        
        if user != nil{
            c.JSON(http.StatusBadRequest,gin.H{
                "message":"invalid",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":"Login succesfull",
            "user":user,
        })
    })
    r.Run(":8080")
}
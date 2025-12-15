package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        
        user,err := c.Cookie("username")
        
        if err != nil{
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":"Invalid",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message":"succes",
        })
        
        c.Next()
        
    }
}

func main() {

	r := gin.Default()
    
    r.POST("/login",func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username == "admin" && password == "1234"{
            c.SetCookie("username", "user", 3600, "/", "localhost",false, true)
            
            c.JSON(http.StatusOK, gin.H{
                "message":"Logged in",
            })
            return
        }
        
        
    })
    
    r.GET("/dashboard", Middleware(),func(c *gin.Context) {
        c.JSON(http.StatusOK,gin.H{
            "message":"Welcome to dashboard",
        })
    })
    r.Run(":8080")
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        session,err := c.Cookie("session")
        
        if err != nil || session != "logged_in"{
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":"Unauthorized",
            })
            c.Abort()
            return 
        }
        c.Next()
    }
}


func main() {
	r := gin.Default()
    
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "Message":"Home",
        })
    })
    
    r.POST("/login", func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username != "admin" && password != "1234"{
            c.JSON(http.StatusUnauthorized, gin.H{
                "message":"invalid credentials",
            })
            return
        }
        
        c.SetCookie("session","logged_in",3600, "/","localhost", false, true)
        
        c.JSON(http.StatusOK,gin.H{
            "message":"Logged in",
        })
    })
    
    r.GET("/dashboard",AuthMiddleware(),func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to dashboard")
    })
    
    r.GET("/logout",func(c *gin.Context) {
        c.SetCookie("session", "" ,-1, "/","localhost",false, true)
        
        c.JSON(http.StatusOK, gin.H{
            "Message":"Logout succesfull",
        })
    })
    r.Run(":8080")
}
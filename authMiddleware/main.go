package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username !=  "admin" && password != "1234"{
            c.JSON(http.StatusOK,gin.H{
                "error":"Invalid credentials",
            })
            c.Abort()
            return     
        }
        c.Next()
    }
}

func main(){
    router := gin.Default()
    
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK,"shinas muhammed k")
    })
    
    router.POST("/dashboard",AuthMiddleware(),func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to dashboard")
    })
    
    router.Run(":8080")
}
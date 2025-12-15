package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        method := c.Request.Method
        path := c.Request.URL.Path
        ip := c.ClientIP()
        
        println("LOG:",method,path,"requested Ip:",ip)
        c.Next()
    }
}

func main() {
    
    r := gin.Default()
    
    r.Use(LoggingMiddleware())
    
    r.POST("/login",func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "Message":"Logged in",
        })
    })
    
    r.GET("/logout", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "Message":"Logged out",
        })
    })
    r.Run(":8080")
}
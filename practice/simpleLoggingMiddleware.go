package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleLogger() gin.HandlerFunc{
    return func(c *gin.Context) {
        method := c.Request.Method
        path := c.Request.URL.Path
        ip := c.ClientIP()
        
        println("LOG:",method,path,"Requested IP:",ip)
        c.Next()
    }
}

func main() {
    r:= gin.Default()
    
    r.Use(SimpleLogger())
    
    r.GET("/login",func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message":"Success",
        })
    })
    r.Run(":8080")
}
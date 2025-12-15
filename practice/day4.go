package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginLogoutMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        method := c.Request.Method
        path := c.Request.URL.Path
        ip := c.ClientIP()
        
        println("LOG:",method,path,"Requested Ip : ",ip)
        c.Next()
    }
}


func AuthMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        
        user,err := c.Cookie("user")
        
        if err != nil || user == ""{
            c.JSON(http.StatusUnauthorized, gin.H{
                "error":"Unauthorized pls login first",
            })
            c.Abort()
            return 
        }
        c.Next()
    }
}

var requestCount = make(map[string]int)
var resetTime = time.Now().Add(1 * time.Minute)


func RateLimiter() gin.HandlerFunc{
    return func(c *gin.Context) {
        ip := c.ClientIP()
        
        if time.Now().After(resetTime){
            requestCount = make(map[string]int)
            resetTime = time.Now().Add(1 * time.Minute)
        }
        
        requestCount[ip]++
        
        if requestCount[ip] >5 {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error":"Rate limit exceeded. Try again after 1 minute",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}

func kwejbkmain() {
    r := gin.Default()
    r.Use(RateLimiter())
    r.Use(LoginLogoutMiddleware())
    
    r.POST("/login",func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username == "admin" && password == "1234"{
            c.SetCookie("user",username, 3600, "/","", false,true)
            
            c.JSON(http.StatusOK,gin.H{
                "message":"Login successfull",
            })
            return
        }
        
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":"Invalid credentials",
        })
    })
    
    
    r.GET("/logout",func(c *gin.Context) {
        c.SetCookie("user", "", -1,"/","", false,true)
        
        c.JSON(http.StatusOK,gin.H{
            "Message":"Logout successfully",
        })
    })
    
    
    r.GET("/dashboard",func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message":"Welcome to Dashboard",
        })
    })
    
    
    r.Run(":8080")
}
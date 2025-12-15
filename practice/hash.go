package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var storeHashedPassword string

func main(){
    r := gin.Default()
    
    r.POST("/login", func(c *gin.Context) {
        
        // username := c.PostForm("username")
        // password := c.PostForm("password")
        
        
        hashed,_ := bcrypt.GenerateFromPassword([]byte("admin123"),bcrypt.DefaultCost)
        
        storeHashedPassword = string(hashed)
        
        var input LoginReq
        
        if err := c.ShouldBindJSON(&input); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "message":"invalid",
            })
            return
        }
        
        err := bcrypt.CompareHashAndPassword([]byte(storeHashedPassword), []byte(input.Password))
        
        if err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "message":"incorrect",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":"logged in",
            "user":input.Username,
        })
                
    })
    
    r.Run(":8080")
}
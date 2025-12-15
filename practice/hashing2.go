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

var storedHashedPassword string

func main() {

	r := gin.Default()
    
    hashed,_ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
    
    storedHashedPassword = string(hashed)
    
    r.POST("/login",func(c *gin.Context) {
        var input LoginReq
        
        if err := c.ShouldBindJSON(&input); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"invalid req",
            })
            return
        }
        
        if input.Username == ""{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"invalid username",
            })
            return
        }
        
        if input.Password == ""{
            c.JSON(http.StatusBadRequest, gin.H{
                "errror":"Invalid username",
            })
            return
        }
        
        err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(input.Password))
        
        if err != nil{
            c.JSON(http.StatusUnauthorized, gin.H{
                "errror":"error",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message":"user logged in",
            "user":input.Username,
        })
    })
    r.Run(":8080")
}
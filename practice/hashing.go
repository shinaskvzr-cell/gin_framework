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
    
    hashed,_ := bcrypt.GenerateFromPassword([]byte("admin123"),bcrypt.DefaultCost)
    
    storedHashedPassword := string(hashed)
    
    r.POST("/login",func(c *gin.Context) {
        var input LoginReq
        
        if err := c.ShouldBindJSON(&input); err!=nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Invalid credential",
            })
            return
        }
        if input.Username == ""{
            c.JSON(http.StatusBadRequest,gin.H{
                "error":"Invalid username",
            })
            return
        }
        if input.Password == ""{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Invalid password",
            })
            return
        }
        
        err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword),[]byte(input.Password))
        
        if err != nil{
            c.JSON(http.StatusUnauthorized,gin.H{
                "error":"password mismatch",
            })
            return
        }
        
        c.JSON(http.StatusOK,gin.H{
            "message":"User logged in",
            "user":input.Username,
        })
    })
    r.Run(":8080")
}

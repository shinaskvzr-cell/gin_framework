package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=4"`
}

func main() {
	r := gin.Default()
    
    
        r.POST("/login",func(c *gin.Context) {
            username := c.PostForm("username")
            password := c.PostForm("password")
            pass:="1234"
            
            hash,_:=bcrypt.GenerateFromPassword([]byte(pass),14)
            hash2 := string(hash)
            fmt.Println(hash2)
            if username == "admin" {
                
                err:=bcrypt.CompareHashAndPassword([]byte(hash2), []byte(password))
                if err==nil {
                    
                    c.JSON(http.StatusOK,gin.H{
                        "success":"login succesfull",
                    })
                }
            }
        })
        r.Run(":8080")
    
}
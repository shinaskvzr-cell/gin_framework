package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required,min=3"`
	Password string `form:"password" binding:"required,min=4"`
}

func main() {
	router := gin.Default()
    
    router.POST("/login", func(c *gin.Context) {
        var input LoginForm
        
        if err := c.ShouldBind(&input) ; err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":err.Error(),
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":"Login success",
            "username":input.Username,
            "userPassword":input.Password,
        })
    })
    router.Run(":8080")
}
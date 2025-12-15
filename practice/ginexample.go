package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string
}

func main() {
	r := gin.Default()
    
    r.POST("/login", func(c *gin.Context) {
        var input Login
        
        if err := c.ShouldBindJSON(&input);err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "message":"Invalid",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":input,
        })
    })
    r.Run(":8080")
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	router := gin.Default()
    api := router.Group("/api")
    api.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK,users)
    })
    
    api.POST("/users", func(c *gin.Context) {
        var newUser User
        if err := c.ShouldBindJSON(&newUser);err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Invalid JSON",
            })
            return
        }
        users = append(users, newUser)
        c.JSON(http.StatusOK, gin.H{
            "status":"User added",
            "user":newUser,
        })
    })
    router.Run(":8080")
}
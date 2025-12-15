package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserProps struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var userList = []UserProps{
	{ID: 1, Name: "Shinas"},
}

func main() {
	r := gin.Default()
    api := r.Group("/api")
    
    api.GET("/users",func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "users":userList,
        })
    })
    
    api.POST("/user",func(c *gin.Context) {
        var newUser UserProps
        
        if err := c.ShouldBindJSON(&newUser); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":err.Error(),
            })
            return
        }
        newUser.ID = len(userList)+1
        userList = append(userList, newUser)
        
        c.JSON(http.StatusOK, gin.H{
            "message":"user added",
            "user":newUser,
        })
    })
    r.Run(":8080")
}
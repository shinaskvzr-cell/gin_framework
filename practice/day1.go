package main

import (
	 "net/http"

	 "github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "shinas", Email: "shinas@gmail.com"},
}

func main() {
	r := gin.Default()
    
    api := r.Group("/api")
    
    api.GET("/users",func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "User":users,
        })
    })
    
   api.POST("/users",func(c *gin.Context) {
        var newUser User
        
        if err := c.ShouldBindJSON(&newUser); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Invalid input",
            })
            return
        }
        newUser.ID = len(users)+1
        users = append(users, newUser)
        
        c.JSON(http.StatusOK,gin.H{
            "message":"user added",
            "user":newUser,
        })
   })
    
    r.Run(":8080")
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Loginrequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	r := gin.Default()
    
    r.POST("/login",func(c *gin.Context) {
        var input Loginrequest
        
        if err := c.ShouldBindJSON(&input) ; err != nil{
            c.JSON(http.StatusBadRequest,gin.H{
                "error":"Invalid JSON format",
            })
            return
        }
        
        if input.Username == ""{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Username is required",
            })
            return
        }
        
        if input.Password == ""{
            c.JSON(http.StatusBadRequest, gin.H{
                "error":"Password is required",
            })
            return
        }
        if len(input.Password)<6{
            c.JSON(http.StatusBadRequest,gin.H{
                "error":"password must be atleast 6 characters",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message":"Login successfull",
            "username":input.Username,
            "password":input.Password,
        })
    })
    r.Run(":8080")
}
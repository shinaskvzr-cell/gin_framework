package main

import "github.com/gin-gonic/gin"

type body struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func main() {
	//create a router
	router := gin.Default()
    var data body

	//define a GET route
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": data,
		})
	})

	//define a POST route
	router.POST("/user", func(c *gin.Context) {
        var data body

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		//send response
		c.JSON(200, gin.H{
			"status": "User recieved",
			"name":   data.Name,
			"age":    data.Age,
		})

	})

	//start server on 8080
	router.Run(":8080")
}

package main

import "github.com/gin-gonic/gin"

func main(){

	r := gin.Default()
	r.GET("/test", func( c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})




	r.Run("127.0.0.1:8080")

}
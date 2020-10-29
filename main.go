package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/gin-gonic/gin"

	"github.com/nurrizkyimani/fabelio/database"
)	

func main(){


	index := database.InitAlgolia()

	

	
	

	fmt.Printf("print")
	
	r := gin.Default()

	r.GET("/search/:keyword", func( c *gin.Context){

		keyword := c.Param("keyword")

		// _ := c.Request.Body


		
		params := []interface{} {
			opt.AttributesToRetrieve("ProductName", "Colours"),	
		}

		res , err := index.Search(keyword, params...)

		if err != nil {
			panic("panic on indexing in get ")
		}


		c.JSON(200, res)

	})
	r.GET("/test", func( c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})



	r.Run("127.0.0.1:8080")

}
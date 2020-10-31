package main

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/gin-gonic/gin"
	"github.com/nurrizkyimani/fabelio/database"
	"github.com/nurrizkyimani/fabelio/model"
)	

func main(){


	index := database.InitAlgolia()

	

	fmt.Printf("print")
	
	r := gin.Default()

	r.GET("/search/:keyword", func( c *gin.Context){

		keyword := c.Param("keyword")

		// _ := c.Request.Body


		// jsonData, err := ioutil.ReadAll(c.Request.Body)
		// if err != nil {
		// 		// Handle error
		// }



		
		params := []interface{} {
			opt.AttributesToRetrieve("ProductName", "Colours"),	
		}

		res , err := index.Search(keyword, params...)
		reshit := res.Hits

		if err != nil {
			panic("panic on indexing in get ")
		}

		
		b, err := json.Marshal(reshit)

		var a[] model.Hit
		err = json.Unmarshal(b, &a)

		if err != nil {
				fmt.Println("error:", err)
		}

		o := a[0].ObjectID

		newUpdate := model.ProductSeenUpdate{
			ObjectID: o ,
			Seen: true,
		}	

		res1, err := index.SaveObject(newUpdate)

		if err != nil {
				fmt.Println("error:", err)
		}
		
		fmt.Println(res1)
		
		c.JSON(200, a[0] )

	})



	r.GET("/test", func( c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})




	r.Run("127.0.0.1:8080")

}
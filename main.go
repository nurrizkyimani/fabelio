package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
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
		decodedValue, err := url.QueryUnescape(keyword)


		index.SetSettings(search.Settings{
			AttributesForFaceting: opt.AttributesForFaceting(
				"Seen", // or "filterOnly(brand)" for filtering purposes only
			),
		})
		
		params := []interface{} {
			opt.AttributesToRetrieve("ProductName", "Colours"),	
			opt.Filters("Seen:FALSE OR Seen:false"),

		}

		res , err := index.Search(decodedValue, params...)
		reshit := res.Hits

		if err != nil {
			panic("panic on indexing in get ")
		}

		//marshaling the hits json;
		b, err := json.Marshal(reshit)

		//unmarshal the b;
		var a[] model.Hit
		err = json.Unmarshal(b, &a)

		if err != nil {
				fmt.Println("error:", err)
		}

		//if len a is not zero;
		if len(a) != 0 {
			o := a[0].ObjectID

			newUpdate := model.ProductSeenUpdate{
				ObjectID: o ,
				Seen: true,
			}	

			res1, err := index.PartialUpdateObject(newUpdate)

			if err != nil {
					fmt.Println("error:", err)
			}
			
			fmt.Println(res1)
		}
		//end of if


		//return if len zero
		
		if len(a) == 0 {
			c.JSON(200, a)
			return
		} else {
			//Return
			c.JSON(200, a[0])
			return
		}
	})



	r.GET("/test", func( c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})


	r.GET("/reload", func( c*gin.Context){


		c.JSON(200, gin.H {
			"message": "Successful Reload",
		})
	})


	r.Run("127.0.0.1:8080")

}
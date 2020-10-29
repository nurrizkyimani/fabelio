package datbase

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
)


func initAlgolia() (*search.Index, *search.Client) {

	e := godotenv.Load()

	if e != nil {
		panic("no env file")
	}

	YourApplicationID := os.Getenv("YourApplicationID")
	YourAdminAPIKey := os.Getenv("YourAdminAPIKey")
	your_index_name := os.Getenv("your_index_name")

	client := search.NewClient(YourApplicationID, YourAdminAPIKey)
	index := client.InitIndex(your_index_name)

	return index, client 
}


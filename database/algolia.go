package database

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/joho/godotenv"
)

//initAlgolia is a fucntion
func InitAlgolia() (*search.Index) {

	e := godotenv.Load()

	if e != nil {
		panic("no env file")
	}

	YourApplicationID := os.Getenv("YourApplicationID")
	YourAdminAPIKey := os.Getenv("YourAdminAPIKey")
	your_index_name := os.Getenv("your_index_name")

	client := search.NewClient(YourApplicationID, YourAdminAPIKey)
	index := client.InitIndex(your_index_name)

	return index
}


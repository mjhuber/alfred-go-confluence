package main

import (
	"fmt"
	"log"
	"os"

	confluence "github.com/virtomize/confluence-go-api"
)

func main() {

	url := os.Getenv("CONFLUENCE_BASEURL")
	username := os.Getenv("CONFLUENCE_USERNAME")
	token := os.Getenv("CONFLUENCE_TOKEN")

	api, err := confluence.NewAPI(url, username, token)
	if err != nil {
		log.Fatal(err)
	}

	// define your Search parameters
	query := confluence.SearchQuery{
		CQL: fmt.Sprintf("type=page and title~'%s'", os.Args[1]),
	}

	// execute search
	result, err := api.Search(query)
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range result.Results {
		fmt.Printf("%s: %s\n", v.Content.Title, v.URL)
	}
}

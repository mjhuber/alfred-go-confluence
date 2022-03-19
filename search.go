package main

import (
	"fmt"
	"log"

	confluence "github.com/virtomize/confluence-go-api"
)

func search(token string, baseurl string, username string, searchQuery string) {
	showUpdateStatus()
	api, err := confluence.NewAPI(fmt.Sprintf("%s/wiki/rest/api", baseurl), username, token)
	if err != nil {
		log.Fatal(err)
	}

	// define your Search parameters
	query := confluence.SearchQuery{
		CQL: fmt.Sprintf("type=page and title~'%s'", searchQuery),
	}

	// execute search
	result, err := api.Search(query)
	if err != nil {
		log.Fatal(err)
	}

	// loop over results
	for _, v := range result.Results {
		wf.NewItem(v.Content.Title).Valid(true).Arg(fmt.Sprintf("%s/wiki%s", baseurl, v.URL)).Subtitle(v.Content.Space.Key)
	}
}

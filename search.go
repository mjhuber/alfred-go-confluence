package main

import (
	"fmt"
	"log"

	confluence "github.com/virtomize/confluence-go-api"
)

func search(cfg *Options, searchQuery string) {
	showUpdateStatus()
	api, err := confluence.NewAPI(fmt.Sprintf("%s/wiki/rest/api", cfg.BaseURL), cfg.Username, cfg.Token)
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
		url := fmt.Sprintf("%s/wiki%s", cfg.BaseURL, v.URL)
		wf.NewItem(v.Content.Title).Valid(true).Arg(url).Subtitle(url)
	}
}

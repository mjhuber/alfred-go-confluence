package main

import (
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	repo  = "mjhuber/alfred-go-confluence"
	query string

	// aw.Workflow is the main API
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func main() {
	wf.Run(run)
}

func run() {
	//wf.NewItem("Hello World!").Valid(true).Var("ARG", "yahoo.com").Subtitle("this is a subtitle")

	errors := []string{}
	token, tFound := wf.Config.Env.Lookup("CONFLUENCE_TOKEN")
	if !tFound {
		errors = append(errors, "CONFLUENCE_TOKEN is not set!")
	}

	baseurl, bFound := wf.Config.Env.Lookup("CONFLUENCE_BASEURL")
	if !bFound {
		errors = append(errors, "CONFLUENCE_BASEURL is not set!")
	}

	username, uFound := wf.Config.Env.Lookup("CONFLUENCE_USERNAME")
	if !uFound {
		errors = append(errors, "CONFLUENCE_USERNAME is not set!")
	}

	if len(errors) > 0 {
		for _, err := range errors {
			wf.WarnEmpty(err, "")
		}
		wf.SendFeedback()
		return
	}

	search(token, baseurl, username, os.Args[1])
	wf.SendFeedback()
}

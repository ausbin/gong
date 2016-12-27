package main

import (
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args)-1 != 1 {
		log.Fatalln("you must pass only the path to a git repository")
	}

	templates, err := templates.NewLoader("templates")

	if err != nil {
		log.Fatalln(err)
	}

	repo, err := models.NewRepo("gong", "a git repository viewer", os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/tree/", http.StripPrefix("/tree", handlers.NewTreeHandler(repo, templates.Get("repo-tree"))))
	// Ideally, the downstream server (nginx, Apache, etc.) would handle
	// requests to /static/ instead, but this is useful for testing.
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	err = http.ListenAndServe(":8050", mux)

	if err != nil {
		log.Fatalln(err)
	}
}

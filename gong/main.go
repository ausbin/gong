package main

import (
	"code.austinjadams.com/gong"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args)-1 != 1 {
		log.Fatalln("you must pass only the path to a git repository")
	}

	templ, err := template.ParseFiles("templates/template.html")

	if err != nil {
		log.Fatalln(err)
	}

	repo, err := gong.NewRepo(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/tree/", http.StripPrefix("/tree", gong.NewTreeHandler(repo, templ)))
	// Ideally, the downstream server (nginx, Apache, etc.) would handle
	// requests to /static/ instead, but this is useful for testing.
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	err = http.ListenAndServe(":8050", mux)

	if err != nil {
		log.Fatalln(err)
	}
}

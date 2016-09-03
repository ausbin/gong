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

	templ, err := template.ParseFiles("template.html")

	if err != nil {
		log.Fatalln(err)
	}

	repo, err := gong.NewRepo(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/tree/", handlerFactory("/tree/", templ, repo))
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	err = http.ListenAndServe(":8050", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

type ctx struct {
	Name, Desc string
	Files      []*gong.RepoFile
}

func handlerFactory(prefix string, templ *template.Template, repo *gong.Repo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir := r.URL.Path[len(prefix):]
		files, err := repo.ListFiles("master", dir)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templ.Execute(w, &ctx{"foo", "a test repository", files})
	})
}

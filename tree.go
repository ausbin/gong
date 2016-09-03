package gong

import (
	"html/template"
	"log"
	"net/http"
)

type TreeHandler struct {
	repo  *Repo
	templ *template.Template
}

func NewTreeHandler(repo *Repo, templ *template.Template) *TreeHandler {
	return &TreeHandler{repo, templ}
}

func (th *TreeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	files, err := th.repo.ListFiles("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = th.templ.Execute(w, NewTreeContext(th.repo, files))
	if err != nil {
		log.Println(err)
	}
}

type TreeContext struct {
	Repo  *Repo
	Files []*RepoFile
}

func NewTreeContext(repo *Repo, files []*RepoFile) *TreeContext {
	return &TreeContext{repo, files}
}

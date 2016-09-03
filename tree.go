package gong

import (
	"log"
	"net/http"
	"html/template"
)

type TreeHandler struct {
	repo *Repo
	templ *template.Template
}

func NewTreeHandler(repo *Repo, templ *template.Template) *TreeHandler {
	return &TreeHandler{repo, templ}
}

func (th *TreeHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	files, err := th.repo.ListFiles("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = th.templ.Execute(w, NewTreeContext("foo", "a test repository", files))
	if err != nil {
		log.Println(err)
	}
}

type TreeContext struct {
	Name, Desc string
	Files      []*RepoFile
}

func NewTreeContext(name, desc string, files []*RepoFile) *TreeContext {
	return &TreeContext{name, desc, files}
}

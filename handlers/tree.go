package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type Tree struct {
	repo  *models.Repo
	templ *template.Template
}

func NewTree(repo *models.Repo, templ *template.Template) *Tree {
	return &Tree{repo, templ}
}

func (th *Tree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	files, err := th.repo.ListFiles("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = th.templ.Execute(w, NewTreeContext(th.repo, path == "/", files))
	if err != nil {
		log.Println(err)
	}
}

type TreeContext struct {
	Repo   *models.Repo
	IsRoot bool
	Files  []models.RepoFile
}

func NewTreeContext(repo *models.Repo, isRoot bool, files []models.RepoFile) *TreeContext {
	return &TreeContext{repo, isRoot, files}
}

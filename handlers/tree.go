package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type TreeHandler struct {
	repo  *models.Repo
	templ *template.Template
}

func NewTreeHandler(repo *models.Repo, templ *template.Template) *TreeHandler {
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
	Repo  *models.Repo
	Files []*models.RepoFile
}

func NewTreeContext(repo *models.Repo, files []*models.RepoFile) *TreeContext {
	return &TreeContext{repo, files}
}

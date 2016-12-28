package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type RepoTree struct {
	repo  *models.Repo
	templ *template.Template
}

func NewRepoTree(repo *models.Repo, templ *template.Template) *RepoTree {
	return &RepoTree{repo, templ}
}

func (th *RepoTree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// XXX Don't hardcode the path here
	path := r.URL.Path[len("/"+th.repo.Name+"/tree"):]

	// Redirect if path does not end in /
	if path[len(path)-1] != '/' {
		http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
		return
	}

	files, err := th.repo.ListFiles("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = th.templ.Execute(w, &repoTreeContext{th.repo, path == "/", path, files})
	if err != nil {
		log.Println(err)
	}
}

type repoTreeContext struct {
	Repo   *models.Repo
	IsRoot bool
	Path   string
	Files  []models.RepoFile
}

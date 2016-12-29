package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoTree struct {
	url   url.Reverser
	repo  *models.Repo
	templ *template.Template
}

func NewRepoTree(url url.Reverser, repo *models.Repo, templ *template.Template) *RepoTree {
	return &RepoTree{url, repo, templ}
}

func (rt *RepoTree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// XXX Don't hardcode the path here
	path := r.URL.Path[len("/"+rt.repo.Name+"/tree"):]

	// Redirect if path does not end in /
	if path[len(path)-1] != '/' {
		http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
		return
	}

	files, err := rt.repo.ListFiles("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := ctx.NewRepoTree(rt.url, rt.repo, path == "/", path, files)
	err = rt.templ.Execute(w, ctx)
	if err != nil {
		log.Println(err)
	}
}

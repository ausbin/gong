package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoRoot struct {
	url   url.Reverser
	repo  *models.Repo
	templ *template.Template
}

func NewRepoRoot(url url.Reverser, repo *models.Repo, templ *template.Template) *RepoRoot {
	return &RepoRoot{url, repo, templ}
}

func (rr *RepoRoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files, err := rr.repo.ListFiles("master", "/")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := ctx.NewRepoRoot(rr.url, rr.repo, files,
		template.HTML("this is the <em>readme</em>"))
	err = rr.templ.Execute(w, ctx)
	if err != nil {
		log.Println(err)
	}
}

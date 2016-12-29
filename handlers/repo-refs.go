package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoRefs struct {
	url   url.Reverser
	repo  *models.Repo
	templ *template.Template
}

func NewRepoRefs(url url.Reverser, repo *models.Repo, templ *template.Template) *RepoRefs {
	return &RepoRefs{url, repo, templ}
}

func (rr *RepoRefs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := rr.templ.Execute(w, ctx.NewRepoRefs(rr.url, rr.repo))

	if err != nil {
		log.Println(err)
	}
}

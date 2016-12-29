package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoLog struct {
	url   url.Reverser
	repo  *models.Repo
	templ *template.Template
}

func NewRepoLog(url url.Reverser, repo *models.Repo, templ *template.Template) *RepoLog {
	return &RepoLog{url, repo, templ}
}

func (rl *RepoLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := rl.templ.Execute(w, ctx.NewRepoLog(rl.url, rl.repo))

	if err != nil {
		log.Println(err)
	}
}

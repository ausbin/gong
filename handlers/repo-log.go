package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoLog struct {
	cfg   *config.Global
	url   url.Reverser
	repo  models.Repo
	templ *template.Template
}

func NewRepoLog(cfg *config.Global, url url.Reverser, repo models.Repo, templ *template.Template) *RepoLog {
	return &RepoLog{cfg, url, repo, templ}
}

func (rl *RepoLog) Serve(w http.ResponseWriter, r *http.Request, info Info) {
	err := rl.templ.Execute(w, ctx.NewRepoLog(rl.cfg, rl.url, rl.repo))

	if err != nil {
		log.Println(err)
	}
}

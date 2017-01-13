package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
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

func (rl *RepoLog) Serve(r Request) {
	err := rl.templ.Execute(r, ctx.NewRepoLog(rl.cfg, rl.url, rl.repo))

	if err != nil {
		r.Error(err)
	}
}

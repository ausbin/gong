package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
)

type RepoRefs struct {
	cfg   *config.Global
	url   url.Reverser
	repo  models.Repo
	templ *template.Template
}

func NewRepoRefs(cfg *config.Global, url url.Reverser, repo models.Repo, templ *template.Template) *RepoRefs {
	return &RepoRefs{cfg, url, repo, templ}
}

func (rr *RepoRefs) Serve(r Request) {
	err := rr.templ.Execute(r, ctx.NewRepoRefs(rr.cfg, rr.url, rr.repo))

	if err != nil {
		r.Error(err)
	}
}

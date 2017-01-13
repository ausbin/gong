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

type List struct {
	cfg   *config.Global
	url   url.Reverser
	repos []models.Repo
	templ *template.Template
}

func NewList(cfg *config.Global, url url.Reverser, repos []models.Repo, templ *template.Template) *List {
	return &List{cfg, url, repos, templ}
}

func (l *List) Serve(w http.ResponseWriter, r *http.Request, info Info) {
	ctx := ctx.NewList(l.cfg, l.url, l.repos)

	err := l.templ.Execute(w, ctx)
	if err != nil {
		log.Println(err)
	}
}

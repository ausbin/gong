package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
)

type List struct {
	cfg      *config.Global
	url      url.Reverser
	repos    []models.Repo
	consumer ctx.Consumer
}

func NewList(cfg *config.Global, url url.Reverser, repos []models.Repo, consumer ctx.Consumer) *List {
	return &List{cfg, url, repos, consumer}
}

func (l *List) Serve(r Request) {
	ctx := ctx.NewList(l.cfg, l.url, l.repos)

	err := l.consumer.Consume(r, ctx)
	if err != nil {
		r.Error(err)
	}
}

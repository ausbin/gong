package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
)

type RepoRefs struct {
	cfg      *config.Global
	url      url.Reverser
	repo     models.Repo
	consumer ctx.Consumer
}

func NewRepoRefs(cfg *config.Global, url url.Reverser, repo models.Repo, consumer ctx.Consumer) *RepoRefs {
	return &RepoRefs{cfg, url, repo, consumer}
}

func (rr *RepoRefs) Serve(r Request) {
	err := rr.consumer.Consume(r, ctx.NewRepoRefs(rr.cfg, rr.url, rr.repo))

	if err != nil {
		r.Error(err)
	}
}

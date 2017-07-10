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
	branch := r.QueryString()["h"]

	if branch == "" {
		branch = rr.repo.DefaultBranch()
	}

	branches, err := rr.repo.Branches()

	if err != nil {
		r.Error(err)
		return
	}

	err = rr.consumer.Consume(r, ctx.NewRepoRefs(rr.cfg, rr.url, rr.repo, branch, branches))

	if err != nil {
		r.Error(err)
	}
}

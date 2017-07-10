package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
)

type RepoLog struct {
	cfg      *config.Global
	url      url.Reverser
	repo     models.Repo
	consumer ctx.Consumer
}

func NewRepoLog(cfg *config.Global, url url.Reverser, repo models.Repo, consumer ctx.Consumer) *RepoLog {
	return &RepoLog{cfg, url, repo, consumer}
}

func (rl *RepoLog) Serve(r Request) {
	branch := r.QueryString()["h"]

	if branch == "" {
		branch = rl.repo.DefaultBranch()
	}

	branches, err := rl.repo.Branches()

	if err != nil {
		r.Error(err)
		return
	}

	err = rl.consumer.Consume(r, ctx.NewRepoLog(rl.cfg, rl.url, rl.repo, branch, branches))

	if err != nil {
		r.Error(err)
	}
}

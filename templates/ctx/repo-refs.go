package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoRefs interface {
	RepoGlobal
}

func NewRepoRefs(cfg *config.Global, url url.Reverser, repo models.Repo, branch string, branches []string) RepoRefs {
	return &repoRefs{NewRepoGlobal(cfg, url, repo, branch, branches)}
}

type repoRefs struct {
	RepoGlobal
}

package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoLog interface {
	RepoGlobal
}

func NewRepoLog(cfg *config.Global, url url.Reverser, repo models.Repo, branch string, branches []string) RepoLog {
	return &repoLog{NewRepoGlobal(cfg, url, repo, branch, branches)}
}

type repoLog struct {
	RepoGlobal
}

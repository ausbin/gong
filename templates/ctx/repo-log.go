package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoLog interface {
	RepoGlobal
}

func NewRepoLog(url url.Reverser, repo *models.Repo) RepoLog {
	return &repoLog{NewRepoGlobal(url, repo)}
}

type repoLog struct {
	RepoGlobal
}

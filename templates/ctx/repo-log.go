package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoLog interface {
	Repo
}

type repoLog struct {
	*repo
}

func NewRepoLog(url url.Reverser, repo *models.Repo) RepoLog {
	return &repoLog{newRepo(url, repo)}
}

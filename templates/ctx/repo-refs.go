package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoRefs interface {
	RepoGlobal
}

func NewRepoRefs(url url.Reverser, repo *models.Repo) RepoRefs {
	return &repoRefs{NewRepoGlobal(url, repo)}
}

type repoRefs struct {
	RepoGlobal
}

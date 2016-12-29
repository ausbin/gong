package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoRefs interface {
	Repo
}

type repoRefs struct {
	*repo
}

func NewRepoRefs(url url.Reverser, repo *models.Repo) RepoRefs {
	return &repoRefs{newRepo(url, repo)}
}

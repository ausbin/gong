package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

// Originally named Repo, but renamed to RepoGlobal to resolve conflict with
// the embedded Repo object and the Repo() method in subclasses
type RepoGlobal interface {
	Global

	Repo() *models.Repo
}

func NewRepoGlobal(url url.Reverser, repo *models.Repo) RepoGlobal {
	return &repoGlobal{NewGlobal(url), repo}
}

type repoGlobal struct {
	Global

	repo *models.Repo
}

func (r *repoGlobal) Repo() *models.Repo { return r.repo }

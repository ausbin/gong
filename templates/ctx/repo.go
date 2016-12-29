package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type Repo interface {
	Global

	Repo() *models.Repo
}

func NewRepo(url url.Reverser, repo *models.Repo) Repo {
	return newRepo(url, repo)
}

type repo struct {
	*global

	repo *models.Repo
}

func (r *repo) Repo() *models.Repo { return r.repo }

func newRepo(url url.Reverser, repo_ *models.Repo) *repo {
	return &repo{newGlobal(url), repo_}
}

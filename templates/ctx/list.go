package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type List interface {
	Global

	Repos() []*models.Repo
}

func NewList(cfg *config.Global, url url.Reverser, repos []*models.Repo) List {
	return &list{NewGlobal(cfg, url), repos}
}

type list struct {
	Global

	repos []*models.Repo
}

func (l *list) Repos() []*models.Repo { return l.repos }

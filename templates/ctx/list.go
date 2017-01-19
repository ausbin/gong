package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type List interface {
	Global

	Repos() []models.Repo
}

func NewList(cfg *config.Global, url url.Reverser, repos []models.Repo) List {
	return &list{NewGlobal(cfg, url), repos}
}

type list struct {
	Global

	repos []models.Repo
}

func (l *list) Repos() []models.Repo { return l.repos }

// XXX Don't
func (l *list) reposEqual(a, b []models.Repo) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}

func (l *list) Equals(other Global) bool {
	otherList, ok := other.(List)

	return ok && l.Global.Equals(other) &&
		l.reposEqual(l.Repos(), otherList.Repos())
}

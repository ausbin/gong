package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

// Originally named Repo, but renamed to RepoGlobal to resolve conflict with
// the embedded Repo object and the Repo() method in subclasses
type RepoGlobal interface {
	Global

	Repo() models.Repo
	Branch() string
	Branches() []string
}

func NewRepoGlobal(cfg *config.Global, url url.Reverser, repo models.Repo, branch string, branches []string) RepoGlobal {
	return &repoGlobal{NewGlobal(cfg, url), repo, branch, branches}
}

type repoGlobal struct {
	Global

	repo     models.Repo
	branch   string
	branches []string
}

func (r *repoGlobal) Repo() models.Repo  { return r.repo }
func (r *repoGlobal) Branch() string     { return r.branch }
func (r *repoGlobal) Branches() []string { return r.branches }

func (r *repoGlobal) Equals(other Global) bool {
	otherGlobal, ok := other.(RepoGlobal)

	return ok && r.Global.Equals(other) &&
		r.Repo() == otherGlobal.Repo()
}

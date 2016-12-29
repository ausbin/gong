package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
)

type RepoRoot interface {
	RepoTree

	Readme() template.HTML
}

func NewRepoRoot(url url.Reverser, repo *models.Repo, files []models.RepoFile,
	readme template.HTML) RepoRoot {
	return newRepoRoot(url, repo, files, readme)
}

type repoRoot struct {
	*repoTree

	readme template.HTML
}

func (r *repoRoot) Readme() template.HTML { return r.readme }

func newRepoRoot(url url.Reverser, repo *models.Repo, files []models.RepoFile,
	readme template.HTML) *repoRoot {
	return &repoRoot{newRepoTree(url, repo, true, "/", files), readme}
}

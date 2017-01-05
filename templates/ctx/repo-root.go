package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
)

type RepoRoot interface {
	RepoTree

	Readme() template.HTML
}

func NewRepoRoot(cfg *config.Global, url url.Reverser, repo *models.Repo, files []models.RepoFile, readme template.HTML) RepoRoot {
	return &repoRoot{NewRepoTree(cfg, url, repo, true, "/", true, files, ""), readme}
}

type repoRoot struct {
	RepoTree

	readme template.HTML
}

func (r *repoRoot) Readme() template.HTML { return r.readme }

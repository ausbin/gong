package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
)

type RepoRoot interface {
	RepoTree

	ReadmePlain() string
	ReadmeHTML() template.HTML
}

func NewRepoRoot(cfg *config.Global, url url.Reverser, repo models.Repo, branch string, branches []string, files []models.RepoFile, readme string, isReadmeHTML bool) RepoRoot {
	return &repoRoot{NewRepoTree(cfg, url, repo, branch, branches, "/", true, false, false, files, ""), readme, isReadmeHTML}
}

type repoRoot struct {
	RepoTree

	readme       string
	isReadmeHTML bool
}

func (r *repoRoot) ReadmePlain() string {
	if r.isReadmeHTML {
		return ""
	} else {
		return r.readme
	}
}

func (r *repoRoot) ReadmeHTML() template.HTML {
	if r.isReadmeHTML {
		return template.HTML(r.readme)
	} else {
		return template.HTML("")
	}
}

func (r *repoRoot) Equals(other Global) bool {
	otherRoot, ok := other.(RepoRoot)

	return ok && r.RepoTree.Equals(other) &&
		r.ReadmePlain() == otherRoot.ReadmePlain() &&
		r.ReadmeHTML() == otherRoot.ReadmeHTML()
}

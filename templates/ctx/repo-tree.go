package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoTree interface {
	Repo

	IsRoot() bool
	Path() string
	Files() []models.RepoFile
}

func NewRepoTree(url url.Reverser, repo *models.Repo, isRoot bool, path string, files []models.RepoFile) RepoTree {
	return newRepoTree(url, repo, isRoot, path, files)
}

type repoTree struct {
	*repo

	isRoot bool
	path   string
	files  []models.RepoFile
}

func (t *repoTree) IsRoot() bool             { return t.isRoot }
func (t *repoTree) Path() string             { return t.path }
func (t *repoTree) Files() []models.RepoFile { return t.files }

func newRepoTree(url url.Reverser, repo *models.Repo, isRoot bool, path string, files []models.RepoFile) *repoTree {
	return &repoTree{newRepo(url, repo), isRoot, path, files}
}

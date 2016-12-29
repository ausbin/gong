package ctx

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoTree interface {
	RepoGlobal

	IsRoot() bool
	Path() string
	Files() []models.RepoFile
}

func NewRepoTree(url url.Reverser, repo *models.Repo, isRoot bool, path string, files []models.RepoFile) RepoTree {
	return &repoTree{NewRepoGlobal(url, repo), isRoot, path, files}
}

type repoTree struct {
	RepoGlobal

	isRoot bool
	path   string
	files  []models.RepoFile
}

func (t *repoTree) IsRoot() bool             { return t.isRoot }
func (t *repoTree) Path() string             { return t.path }
func (t *repoTree) Files() []models.RepoFile { return t.files }

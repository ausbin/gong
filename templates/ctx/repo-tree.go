package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoTree interface {
	RepoGlobal

	IsRoot() bool
	Path() string
	IsListing() bool
	Files() []models.RepoFile
	Blob() string
}

func NewRepoTree(cfg *config.Global, url url.Reverser, repo *models.Repo, isRoot bool, path string, isListing bool, files []models.RepoFile, blob string) RepoTree {
	return &repoTree{NewRepoGlobal(cfg, url, repo), isRoot, path, isListing, files, blob}
}

type repoTree struct {
	RepoGlobal

	isRoot    bool
	path      string
	isListing bool
	files     []models.RepoFile
	blob      string
}

func (t *repoTree) IsRoot() bool             { return t.isRoot }
func (t *repoTree) Path() string             { return t.path }
func (t *repoTree) IsListing() bool          { return t.isListing }
func (t *repoTree) Files() []models.RepoFile { return t.files }
func (t *repoTree) Blob() string             { return t.blob }

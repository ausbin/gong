package url

import (
	"code.austinjadams.com/gong/models"
)

type Reverser interface {
	RepoReverser

	Root() string
	Static(string) string
}

type RepoReverser interface {
	RepoRoot(repo *models.Repo) string
	RepoPlain(repo *models.Repo, path string) string
	RepoTree(repo *models.Repo, path string, isDir bool) string
	RepoLog(repo *models.Repo) string
	RepoRefs(repo *models.Repo) string
}

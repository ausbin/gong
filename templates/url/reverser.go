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
	RepoRoot(repo models.Repo, branch string) string
	RepoPlain(repo models.Repo, branch string, path string) string
	RepoTree(repo models.Repo, branch string, path string, isDir bool) string
	RepoLog(repo models.Repo, branch string) string
	RepoRefs(repo models.Repo, branch string) string
}

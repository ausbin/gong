package url

type Reverser interface {
	Root() string

	RepoRoot(repo string) string
	RepoTree(repo string, path string, isDir bool) string
	RepoLog(repo string) string
	RepoRefs(repo string) string
}

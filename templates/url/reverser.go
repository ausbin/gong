package url

type Reverser interface {
	RepoReverser

	Root() string
	Static(string) string
}

type RepoReverser interface {
	RepoRoot(repo string) string
	RepoPlain(repo string, path string) string
	RepoTree(repo string, path string, isDir bool) string
	RepoLog(repo string) string
	RepoRefs(repo string) string
}

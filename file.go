package gong

type RepoFile struct {
	Name  string
	IsDir bool
}

func NewRepoFile(name string, isDir bool) *RepoFile {
	return &RepoFile{name, isDir}
}

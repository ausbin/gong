package models

import (
	git "github.com/libgit2/git2go"
	"log"
)

type RepoFile interface {
	Name() string
	IsFile() bool
	IsDir() bool
	IsExecutable() bool
	Size() int64
	Equals(RepoFile) bool
}

type repoFile struct {
	// Need a reference to the repository to fetch file sizes
	repo  *git.Repository
	entry *git.TreeEntry
}

func (f repoFile) Name() string {
	return f.entry.Name
}

func (f repoFile) IsFile() bool {
	return f.entry.Type == git.ObjectBlob
}

func (f repoFile) IsDir() bool {
	return f.entry.Type == git.ObjectTree
}

func (f repoFile) IsExecutable() bool {
	return f.entry.Filemode == git.FilemodeBlobExecutable
}

func (f repoFile) Size() int64 {
	blob, err := f.repo.LookupBlob(f.entry.Id)

	if err == nil {
		return blob.Size()
	} else {
		log.Println("Warning: Could not read size of tree entry", f.entry.Id, "error:", err)
		return -1
	}
}

func (f repoFile) Equals(other RepoFile) bool {
	return f.Name() == other.Name() &&
		f.IsFile() == other.IsFile() &&
		f.IsDir() == other.IsDir() &&
		f.IsExecutable() == other.IsExecutable() &&
		f.Size() == other.Size()
}

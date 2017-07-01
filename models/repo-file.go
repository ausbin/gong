package models

import (
	git "github.com/libgit2/git2go"
	"log"
)

type RepoFile interface {
	// Attributes
	Name() string
	IsFile() bool
	IsDir() bool
	IsExecutable() bool
	Size() int64

	// Actions
	Equals(RepoFile) bool
	GetBlob() (string, error)
	GetBlobBytes() ([]byte, error)
	ListFiles() ([]RepoFile, error)
}

type repoFile struct {
	name     string
	id       *git.Oid
	type_    git.ObjectType
	filemode git.Filemode

	// Need a reference to the repository to fetch file sizes
	repo *git.Repository
}

func newRepoFile(repo *git.Repository, name string, id *git.Oid, type_ git.ObjectType, filemode git.Filemode) RepoFile {
	return &repoFile{name, id, type_, filemode, repo}
}

func newRepoFileFromTreeEntry(repo *git.Repository, entry *git.TreeEntry) RepoFile {
	return &repoFile{entry.Name, entry.Id, entry.Type, entry.Filemode, repo}
}

func (f *repoFile) Name() string {
	return f.name
}

func (f *repoFile) IsFile() bool {
	return f.type_ == git.ObjectBlob
}

func (f *repoFile) IsDir() bool {
	return f.type_ == git.ObjectTree
}

func (f *repoFile) IsExecutable() bool {
	return f.filemode == git.FilemodeBlobExecutable
}

func (f *repoFile) Size() int64 {
	blob, err := f.repo.LookupBlob(f.id)

	if err == nil {
		return blob.Size()
	} else {
		log.Println("Warning: Could not read size of tree entry", f.id, "error:", err)
		return -1
	}
}

func (f *repoFile) Equals(other RepoFile) bool {
	return f.Name() == other.Name() &&
		f.IsFile() == other.IsFile() &&
		f.IsDir() == other.IsDir() &&
		f.IsExecutable() == other.IsExecutable() &&
		f.Size() == other.Size()
}

func (f *repoFile) GetBlob() (result string, err error) {
	bytes, err := f.GetBlobBytes()

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func (f *repoFile) GetBlobBytes() (result []byte, err error) {
	blob, err := f.repo.LookupBlob(f.id)

	if err != nil {
		return
	}

	result = blob.Contents()
	return
}

func (f *repoFile) ListFiles() (result []RepoFile, err error) {
	tree, err := f.repo.LookupTree(f.id)

	if err != nil {
		return
	}

	tree.Walk(func(_ string, entry *git.TreeEntry) int {
		result = append(result, newRepoFileFromTreeEntry(f.repo, entry))
		return 1
	})

	return
}

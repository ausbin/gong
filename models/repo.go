// Repo model
//
// Handles invocation of git2go

package models

import (
	git "github.com/libgit2/git2go"
)

type Repo interface {
	// Attributes
	Name() string
	Description() string
	Path() string
	DefaultBranch() string

	// Actions
	Open() error
	Readme(branch string) *RepoReadme
	ListFiles(entry *RepoTreeEntry) (files []RepoFile, err error)
	Find(branch, path string) (rte *RepoTreeEntry, err error)
	GetBlob(entry *RepoTreeEntry) (result string, err error)
	GetBlobBytes(entry *RepoTreeEntry) (result []byte, err error)
}

func NewRepo(name, description, path, defbranch string) Repo {
	return &repo{name, description, path, defbranch, nil}
}

type repo struct {
	name, description string
	path              string
	defaultBranch     string
	repo              *git.Repository
}

func (r *repo) Name() string          { return r.name }
func (r *repo) Description() string   { return r.description }
func (r *repo) Path() string          { return r.path }
func (r *repo) DefaultBranch() string { return r.defaultBranch }

func (r *repo) Open() error {
	var err error
	r.repo, err = git.OpenRepository(r.path)
	return err
}

func (r *repo) ListFiles(entry *RepoTreeEntry) (result []RepoFile, err error) {
	tree, err := entry.obj.AsTree()

	if err != nil {
		return
	}

	tree.Walk(func(_ string, entry *git.TreeEntry) int {
		result = append(result, &repoFile{r.repo, entry})
		return 1
	})

	return
}

func (r *repo) GetBlob(entry *RepoTreeEntry) (result string, err error) {
	bytes, err := r.GetBlobBytes(entry)

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func (r *repo) GetBlobBytes(entry *RepoTreeEntry) (result []byte, err error) {
	blob, err := entry.obj.AsBlob()

	if err != nil {
		return
	}

	result = blob.Contents()
	return
}

func (r *repo) Find(branch, path string) (rte *RepoTreeEntry, err error) {
	ref, err := r.repo.LookupBranch(branch, git.BranchLocal)

	if err != nil {
		return
	}

	commitObj, err := ref.Peel(git.ObjectCommit)

	if err != nil {
		return
	}

	commit, err := commitObj.AsCommit()

	if err != nil {
		return
	}

	tree, err := commit.Tree()

	if err != nil {
		return
	}

	var obj *git.Object

	// If we want the tree at the root of the repository, return because
	// we have it. Otherwise, search the root tree for the tree of the
	// desired directory
	if path == "" || path == "/" {
		obj = &tree.Object
	} else {
		// In case path points to a blob instead of a tree, choose to remove a
		// trailing slash if present
		rightOffset := 0
		if path[len(path)-1] == '/' {
			rightOffset = 1
		}

		// Remove leading slash because git2go doesn't accept it, and possibly
		// the trailing slash as described above
		path = path[1 : len(path)-rightOffset]

		var entry *git.TreeEntry
		entry, err = tree.EntryByPath(path)

		if err != nil {
			return
		}

		obj, err = r.repo.Lookup(entry.Id)

		if err != nil {
			return
		}
	}

	rte = &RepoTreeEntry{obj: obj}

	return
}

type RepoTreeEntry struct {
	obj *git.Object
}

func (rte *RepoTreeEntry) IsDir() bool {
	return rte.obj.Type() == git.ObjectTree
}

const (
	RepoReadmeTypePlain RepoReadmeType = iota
	RepoReadmeTypeMarkdown
)

type RepoReadmeType int
type RepoReadme struct {
	Content []byte
	Type    RepoReadmeType
}

func NewRepoReadme(blob []byte, type_ RepoReadmeType) *RepoReadme {
	return &RepoReadme{blob, type_}
}

func (r *repo) Readme(branch string) *RepoReadme {
	readmeNames := []struct {
		name  string
		type_ RepoReadmeType
	}{
		{"/README.md", RepoReadmeTypeMarkdown},
		{"/README", RepoReadmeTypePlain},
	}

	for _, name := range readmeNames {
		entry, err := r.Find(branch, name.name)

		// XXX Handle errors more carefully. err != nil does not
		//     necessarily mean the file doesn't exist in the tree
		if err == nil {
			blob, err := r.GetBlobBytes(entry)

			if err == nil {
				return NewRepoReadme(blob, name.type_)
			}
		}
	}

	return nil
}

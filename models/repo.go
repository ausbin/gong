// Repo model
//
// Handles invocation of git2go

package models

import (
	git "github.com/libgit2/git2go"
)

type Repo struct {
	Name, Description string
	path              string
	repo              *git.Repository
}

func NewRepo(name, description, path string) (*Repo, error) {
	repo, err := git.OpenRepository(path)

	if err != nil {
		return nil, err
	}

	return &Repo{name, description, path, repo}, nil
}

func (r *Repo) ListFiles(entry *RepoTreeEntry) (result []RepoFile, err error) {
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

func (r *Repo) GetBlob(entry *RepoTreeEntry) (result string, err error) {
	bytes, err := r.GetBlobBytes(entry)

	if err != nil {
		return
	}

	result = string(bytes)
	return
}

func (r *Repo) GetBlobBytes(entry *RepoTreeEntry) (result []byte, err error) {
	blob, err := entry.obj.AsBlob()

	if err != nil {
		return
	}

	result = blob.Contents()
	return
}

func (r *Repo) Find(branch, path string) (rte *RepoTreeEntry, err error) {
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

func (r *Repo) Readme(branch string) *RepoReadme {
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

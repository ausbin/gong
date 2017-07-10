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
	Branches() ([]string, error)

	// Actions
	Open() error
	Readme(branch string) *RepoReadme
	Find(branch, path string) (RepoFile, error)
	Equals(other Repo) bool
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

func (r *repo) Branches() (branches []string, err error) {
	iterator, err := r.repo.NewBranchIterator(git.BranchLocal)

	if err != nil {
		return
	}

	err = iterator.ForEach(func(branch *git.Branch, _ git.BranchType) (err error) {
		name, err := branch.Name()

		if err != nil {
			return
		}

		branches = append(branches, name)
		return
	})

	iterator.Free()

	// On error, return a nil slice
	if err != nil {
		branches = nil
	}

	return
}

func (r *repo) Find(branch, path string) (f RepoFile, err error) {
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

	// If we want the tree at the root of the repository, return because
	// we have it. Otherwise, search the root tree for the tree of the
	// desired directory
	if path == "" || path == "/" {
		// Root of the tree does not have a tree entry, so fake one
		f = newRepoFile(r.repo, "/", tree.Id(), git.ObjectTree, git.FilemodeTree)
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

		f = newRepoFileFromTreeEntry(r.repo, entry)
	}

	return
}

func (r *repo) Equals(other Repo) bool {
	return r.Name() == other.Name() &&
		r.Description() == r.Description() &&
		r.Path() == r.Path() &&
		r.DefaultBranch() == r.DefaultBranch()
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
		file, err := r.Find(branch, name.name)

		// XXX Handle errors more carefully. err != nil does not
		//     necessarily mean the file doesn't exist in the tree
		if err == nil {
			blob, err := file.GetBlobBytes()

			if err == nil {
				return NewRepoReadme(blob, name.type_)
			}
		}
	}

	return nil
}

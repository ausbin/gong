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

func (r *Repo) ListFiles(branch, dir string) (result []RepoFile, err error) {
	tree, err := r.tree(branch, dir)

	if err != nil {
		return
	}

	tree.Walk(func(_ string, entry *git.TreeEntry) int {
		result = append(result, &repoFile{r.repo, entry})
		return 1
	})

	return
}

func (r *Repo) tree(branch, dir string) (tree *git.Tree, err error) {
	ref, err := r.repo.LookupBranch(branch, git.BranchLocal)

	if err != nil {
		return
	}

	obj, err := ref.Peel(git.ObjectCommit)

	if err != nil {
		return
	}

	commit, err := obj.AsCommit()

	if err != nil {
		return
	}

	tree, err = commit.Tree()

	if err != nil {
		return
	}

	// If we want the tree at the root of the repository, return because
	// we have it. Otherwise, search the root tree for the tree of the
	// desired directory
	if dir != "" && dir != "/" {
		// Remove leading slash because git2go doesn't accept it
		dir = dir[1:]

		var entry *git.TreeEntry
		entry, err = tree.EntryByPath(dir)

		if err != nil {
			return
		}

		tree, err = r.repo.LookupTree(entry.Id)

		if err != nil {
			return
		}
	}

	return
}

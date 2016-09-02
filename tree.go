package main

import (
	git "github.com/libgit2/git2go"
	"log"
)

func ls(repo *git.Repository, branch, dir string) (result []*ctxfile, err error) {
	tree, err := tree(repo, branch, dir)

	if err != nil {
		return
	}

	tree.Walk(func(_ string, entry *git.TreeEntry) int {
		result = append(result, &ctxfile{entry.Name, entry.Type == git.ObjectTree})
		return 1
	})

	return
}

func tree(repo *git.Repository, branch, dir string) (tree *git.Tree, err error) {
	ref, err := repo.LookupBranch(branch, git.BranchLocal)

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

	if dir != "" && dir != "/" {
		log.Println(dir)
		var entry *git.TreeEntry
		entry, err = tree.EntryByPath(dir)

		if err != nil {
			return
		}

		tree, err = repo.LookupTree(entry.Id)

		if err != nil {
			return
		}
	}

	return
}

package main

import (
	"os"
	"log"
	"net/http"
	"html/template"
	git "github.com/libgit2/git2go"
)

func main() {
	if len(os.Args)-1 != 1 {
		log.Fatalln("you must pass only the path to a git repository")
	}

	templ, err := template.ParseFiles("template.html")

	if err != nil {
		log.Fatalln(err)
	}

	repo, err := git.OpenRepository(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/tree/", handlerFactory("/tree/", templ, repo))
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	err = http.ListenAndServe(":8050", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

type ctx struct {
	Name, Desc string
	Files []*ctxfile
}

type ctxfile struct {
	Name string
	IsDir bool
}

func handlerFactory(prefix string, templ *template.Template, repo *git.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir := r.URL.Path[len(prefix):]
		files, err := ls(repo, "master", dir)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templ.Execute(w, &ctx{"foo", "a test repository", files});
	})
}

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

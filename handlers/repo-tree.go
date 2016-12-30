package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
	"log"
	"net/http"
)

type RepoTree struct {
	url   url.Reverser
	repo  *models.Repo
	templ *template.Template
}

func NewRepoTree(url url.Reverser, repo *models.Repo, templ *template.Template) *RepoTree {
	return &RepoTree{url, repo, templ}
}

func (rt *RepoTree) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The -1 leaves us with a leading slash in the resulting path — so an
	// absolute path
	path := r.URL.Path[len(rt.url.RepoTree(rt.repo.Name, "/", true))-1:]

	entry, err := rt.repo.Find("master", path)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var files []models.RepoFile
	var blob string

	endsWithSlash := r.URL.Path[len(r.URL.Path)-1] == '/'

	if entry.IsDir() {
		// Since this is a directory, redirect if path does not end in /
		if !endsWithSlash {
			http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
			return
		}

		files, err = rt.repo.ListFiles(entry)
	} else {
		// Since this is NOT a directory, redirect if path ends in /
		if endsWithSlash {
			http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
			return
		}

		blob, err = rt.repo.GetBlob(entry)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := ctx.NewRepoTree(rt.url, rt.repo, path == "/", path, entry.IsDir(), files, blob)
	err = rt.templ.Execute(w, ctx)
	if err != nil {
		log.Println(err)
	}
}

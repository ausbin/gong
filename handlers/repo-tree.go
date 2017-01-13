package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"html/template"
)

type RepoTree struct {
	cfg   *config.Global
	url   url.Reverser
	repo  models.Repo
	templ *template.Template
}

func NewRepoTree(cfg *config.Global, url url.Reverser, repo models.Repo, templ *template.Template) *RepoTree {
	return &RepoTree{cfg, url, repo, templ}
}

func (rt *RepoTree) Serve(r Request) {
	path := r.Subtree()
	entry, err := rt.repo.Find(rt.repo.DefaultBranch(), path)

	if err != nil {
		r.Error(err)
		return
	}

	var files []models.RepoFile
	var blob string

	endsWithSlash := r.Path()[len(r.Path())-1] == '/'

	if entry.IsDir() {
		// Since this is a directory, redirect if path does not end in /
		if !endsWithSlash {
			r.Redirect(r.Path() + "/")
			return
		}

		files, err = rt.repo.ListFiles(entry)
	} else {
		// Since this is NOT a directory, redirect if path ends in /
		if endsWithSlash {
			r.Redirect(r.Path()[:len(r.Path())-1])
			return
		}

		blob, err = rt.repo.GetBlob(entry)
	}

	if err == nil {
		ctx := ctx.NewRepoTree(rt.cfg, rt.url, rt.repo, path, entry.IsDir(), files, blob)
		err = rt.templ.Execute(r, ctx)
	}

	if err != nil {
		r.Error(err)
	}
}

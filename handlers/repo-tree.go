package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
)

type RepoTree struct {
	cfg      *config.Global
	url      url.Reverser
	repo     models.Repo
	consumer ctx.Consumer
}

func NewRepoTree(cfg *config.Global, url url.Reverser, repo models.Repo, consumer ctx.Consumer) *RepoTree {
	return &RepoTree{cfg, url, repo, consumer}
}

func (rt *RepoTree) Serve(r Request) {
	path := r.Subtree()
	file, err := rt.repo.Find(rt.repo.DefaultBranch(), path)

	if err != nil {
		r.Error(err)
		return
	}

	var files []models.RepoFile
	var blob string

	endsWithSlash := r.Path()[len(r.Path())-1] == '/'

	if file.IsDir() {
		// Since this is a directory, redirect if path does not end in /
		if !endsWithSlash {
			r.Redirect(r.Path() + "/")
			return
		}

		files, err = file.ListFiles()
	} else {
		// Since this is NOT a directory, redirect if path ends in /
		if endsWithSlash {
			r.Redirect(r.Path()[:len(r.Path())-1])
			return
		}

		blob, err = file.GetBlob()
	}

	if err == nil {
		ctx := ctx.NewRepoTree(rt.cfg, rt.url, rt.repo, path, file.IsDir(), files, blob)
		err = rt.consumer.Consume(r, ctx)
	}

	if err != nil {
		r.Error(err)
	}
}

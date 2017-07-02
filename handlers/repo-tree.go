package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"net/http"
	"strings"
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
	var isBinary, isImage bool

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

		blob_bytes, err := file.GetBlobBytes()

		if err == nil {
			// Sniff the blob's mime type
			mime_type := http.DetectContentType(blob_bytes)
			isBinary = !strings.HasPrefix(mime_type, "text/")
			isImage = strings.HasPrefix(mime_type, "image/")

			blob = string(blob_bytes)
		}
	}

	if err == nil {
		ctx := ctx.NewRepoTree(rt.cfg, rt.url, rt.repo, path, file.IsDir(), isBinary, isImage, files, blob)
		err = rt.consumer.Consume(r, ctx)
	}

	if err != nil {
		r.Error(err)
	}
}

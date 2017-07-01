package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoPlain struct {
	url  url.Reverser
	repo models.Repo
}

func NewRepoPlain(url url.Reverser, repo models.Repo) *RepoPlain {
	return &RepoPlain{url, repo}
}

func (rp *RepoPlain) Serve(r Request) {
	if r.Path()[len(r.Path())-1] == '/' {
		r.Redirect(r.Path()[:len(r.Path())-1])
		return
	}

	path := r.Subtree()
	file, err := rp.repo.Find(rp.repo.DefaultBranch(), path)

	var blob []byte
	if err == nil {
		blob, err = file.GetBlobBytes()
	}

	if err == nil {
		// net/http will call DetectContentType() for us since we didn't set
		// the Content-Type already
		_, err = r.Write(blob)
	}

	if err != nil {
		r.Error(err)
	}
}

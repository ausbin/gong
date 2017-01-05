package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
	"log"
	"net/http"
)

type RepoPlain struct {
	url  url.Reverser
	repo *models.Repo
}

func NewRepoPlain(url url.Reverser, repo *models.Repo) *RepoPlain {
	return &RepoPlain{url, repo}
}

func (rp *RepoPlain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len(r.URL.Path)-1] == '/' {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
		return
	}

	// The -1 leaves us with a leading slash in the resulting path — so an
	// absolute path
	path := r.URL.Path[len(rp.url.RepoPlain(rp.repo.Name, "/"))-1:]

	entry, err := rp.repo.Find("master", path)

	var blob []byte
	if err == nil {
		blob, err = rp.repo.GetBlobBytes(entry)
	}

	if err == nil {
		// net/http will call DetectContentType() for us since we didn't set
		// the Content-Type already
		_, err = w.Write(blob)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
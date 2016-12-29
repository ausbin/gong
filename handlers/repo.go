package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates"
	"code.austinjadams.com/gong/templates/url"
	"net/http"
)

type Repo struct {
	url       url.Reverser
	repo      *models.Repo
	templates templates.Loader
}

func NewRepo(url url.Reverser, repo *models.Repo, templates templates.Loader) *Repo {
	return &Repo{url, repo, templates}
}

func (r *Repo) ConfigureMux(mux *http.ServeMux) {
	root := "/" + r.repo.Name + "/"

	mux.Handle(root, NewRepoRoot(r.url, r.repo, r.templates.Get("repo-root")))
	mux.Handle(root+"tree/", NewRepoTree(r.url, r.repo, r.templates.Get("repo-tree")))
	mux.Handle(root+"log/", NewRepoLog(r.url, r.repo, r.templates.Get("repo-log")))
	mux.Handle(root+"refs/", NewRepoRefs(r.url, r.repo, r.templates.Get("repo-refs")))
}

// XXX Move this out of here
type RepoReverser struct{}

func (r *RepoReverser) RepoRoot(repo string) string {
	return "/" + repo + "/"
}

func (r *RepoReverser) RepoTree(repo string, path string, isDir bool) string {
	result := r.RepoRoot(repo) + "tree" + path
	hasSlash := result[len(result)-1] == '/'

	// Remove/add slash as needed
	if hasSlash && !isDir {
		result = result[:len(path)-1]
	} else if !hasSlash && isDir {
		result += "/"
	}

	return result
}

func (r *RepoReverser) RepoLog(repo string) string {
	return r.RepoRoot(repo) + "log/"
}

func (r *RepoReverser) RepoRefs(repo string) string {
	return r.RepoRoot(repo) + "refs/"
}

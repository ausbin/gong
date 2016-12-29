package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates"
	"net/http"
)

type Repo struct {
	repo      *models.Repo
	templates templates.Loader
}

func NewRepo(repo *models.Repo, templates templates.Loader) *Repo {
	return &Repo{repo, templates}
}

func (r *Repo) ConfigureMux(mux *http.ServeMux) {
	url := &tmpReverser{}
	root := "/" + r.repo.Name + "/"

	mux.Handle(root, NewRepoRoot(url, r.repo, r.templates.Get("repo-root")))
	mux.Handle(root+"tree/", NewRepoTree(url, r.repo, r.templates.Get("repo-tree")))
	mux.Handle(root+"log/", NewRepoLog(url, r.repo, r.templates.Get("repo-log")))
	mux.Handle(root+"refs/", NewRepoRefs(url, r.repo, r.templates.Get("repo-refs")))
}

type tmpReverser struct{}

func (r *tmpReverser) Root() string                                         { return "" }
func (r *tmpReverser) RepoRoot(repo string) string                          { return "" }
func (r *tmpReverser) RepoTree(repo string, path string, isDir bool) string { return "" }
func (r *tmpReverser) RepoLog(repo string) string                           { return "" }
func (r *tmpReverser) RepoRefs(repo string) string                          { return "" }

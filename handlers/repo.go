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
	root := "/" + r.repo.Name + "/"

	mux.Handle(root, NewRepoRoot(r.repo, r.templates.Get("repo-root")))
	mux.Handle(root+"tree/", http.StripPrefix(root+"tree", NewRepoTree(r.repo, r.templates.Get("repo-tree"))))
	mux.Handle(root+"log/", NewRepoLog(r.repo, r.templates.Get("repo-log")))
	mux.Handle(root+"refs/", NewRepoRefs(r.repo, r.templates.Get("repo-refs")))
}

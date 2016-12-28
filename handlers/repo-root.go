package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type RepoRoot struct {
	repo  *models.Repo
	templ *template.Template
}

func NewRepoRoot(repo *models.Repo, templ *template.Template) *RepoRoot {
	return &RepoRoot{repo, templ}
}

func (th *RepoRoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files, err := th.repo.ListFiles("master", "/")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = th.templ.Execute(w, &repoRootContext{th.repo, true, "/", files, template.HTML("this is the <em>readme</em>")})
	if err != nil {
		log.Println(err)
	}
}

type repoRootContext struct {
	Repo   *models.Repo
	IsRoot bool
	Path   string
	Files  []models.RepoFile
	Readme template.HTML
}

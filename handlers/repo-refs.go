package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type RepoRefs struct {
	repo  *models.Repo
	templ *template.Template
}

func NewRepoRefs(repo *models.Repo, templ *template.Template) *RepoRefs {
	return &RepoRefs{repo, templ}
}

func (th *RepoRefs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := th.templ.Execute(w, &repoRefsContext{th.repo})

	if err != nil {
		log.Println(err)
	}
}

type repoRefsContext struct {
	Repo *models.Repo
}

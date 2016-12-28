package handlers

import (
	"code.austinjadams.com/gong/models"
	"html/template"
	"log"
	"net/http"
)

type RepoLog struct {
	repo  *models.Repo
	templ *template.Template
}

func NewRepoLog(repo *models.Repo, templ *template.Template) *RepoLog {
	return &RepoLog{repo, templ}
}

func (th *RepoLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := th.templ.Execute(w, &repoLogContext{th.repo})

	if err != nil {
		log.Println(err)
	}
}

type repoLogContext struct {
	Repo *models.Repo
}

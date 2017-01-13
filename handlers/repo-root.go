package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"github.com/russross/blackfriday"
	"html/template"
)

type RepoRoot struct {
	cfg   *config.Global
	url   url.Reverser
	repo  models.Repo
	templ *template.Template
}

func NewRepoRoot(cfg *config.Global, url url.Reverser, repo models.Repo, templ *template.Template) *RepoRoot {
	return &RepoRoot{cfg, url, repo, templ}
}

func (rr *RepoRoot) Serve(r Request) {
	entry, err := rr.repo.Find(rr.repo.DefaultBranch(), "/")

	var files []models.RepoFile
	if err == nil {
		files, err = rr.repo.ListFiles(entry)
	}

	if err == nil {
		readme, isReadmeHTML := rr.Readme()
		ctx := ctx.NewRepoRoot(rr.cfg, rr.url, rr.repo, files, readme,
			isReadmeHTML)
		err = rr.templ.Execute(r, ctx)
	}

	if err != nil {
		r.Error(err)
	}
}

func (rr *RepoRoot) Readme() (content string, isReadmeHTML bool) {
	readme := rr.repo.Readme(rr.repo.DefaultBranch())

	if readme != nil {
		if readme.Type == models.RepoReadmeTypeMarkdown {
			return string(blackfriday.MarkdownBasic(readme.Content)), true
		} else {
			// RepoReadmeTypePlain
			return string(readme.Content), false
		}
	}

	// No readme was found
	return "", false
}

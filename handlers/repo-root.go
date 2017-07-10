package handlers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"code.austinjadams.com/gong/templates/url"
	"github.com/russross/blackfriday"
)

type RepoRoot struct {
	cfg      *config.Global
	url      url.Reverser
	repo     models.Repo
	consumer ctx.Consumer
}

func NewRepoRoot(cfg *config.Global, url url.Reverser, repo models.Repo, consumer ctx.Consumer) *RepoRoot {
	return &RepoRoot{cfg, url, repo, consumer}
}

func (rr *RepoRoot) Serve(r Request) {
	branch := r.QueryString()["h"]

	if branch == "" {
		branch = rr.repo.DefaultBranch()
	}

	branches, err := rr.repo.Branches()

	if err != nil {
		r.Error(err)
		return
	}

	file, err := rr.repo.Find(branch, "/")

	var files []models.RepoFile
	if err == nil {
		files, err = file.ListFiles()
	}

	if err == nil {
		readme, isReadmeHTML := rr.Readme(branch)
		ctx := ctx.NewRepoRoot(rr.cfg, rr.url, rr.repo, branch, branches,
			files, readme, isReadmeHTML)
		err = rr.consumer.Consume(r, ctx)
	}

	if err != nil {
		r.Error(err)
	}
}

func (rr *RepoRoot) Readme(branch string) (content string, isReadmeHTML bool) {
	readme := rr.repo.Readme(branch)

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

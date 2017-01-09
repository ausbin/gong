package routers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates"
	"code.austinjadams.com/gong/templates/url"
)

type Repo struct {
	cfg       *config.Global
	url       url.Reverser
	repo      *models.Repo
	templates templates.Loader
}

func NewRepo(cfg *config.Global, url url.Reverser, repo *models.Repo, templates templates.Loader) SubRouter {
	return &Repo{cfg, url, repo, templates}
}

func (r *Repo) ConfigureRouter(superRouter Router) {
	superRouter.Handle(r.url.RepoRoot(r.repo), false,
		handlers.NewRepoRoot(r.cfg, r.url, r.repo, r.templates.Get("repo-root")))
	superRouter.Handle(r.url.RepoPlain(r.repo, "/"), true,
		handlers.NewRepoPlain(r.url, r.repo))
	superRouter.Handle(r.url.RepoTree(r.repo, "/", true), true,
		handlers.NewRepoTree(r.cfg, r.url, r.repo, r.templates.Get("repo-tree")))
	superRouter.Handle(r.url.RepoLog(r.repo), false,
		handlers.NewRepoLog(r.cfg, r.url, r.repo, r.templates.Get("repo-log")))
	superRouter.Handle(r.url.RepoRefs(r.repo), false,
		handlers.NewRepoRefs(r.cfg, r.url, r.repo, r.templates.Get("repo-refs")))
}

type repoReverser struct {
	repoPrefix string
}

func NewRepoReverser(repoPrefix string) url.RepoReverser {
	return &repoReverser{repoPrefix}
}

func (r *repoReverser) RepoRoot(repo *models.Repo) string {
	return r.repoPrefix + "/" + repo.Name + "/"
}

func (r *repoReverser) RepoPlain(repo *models.Repo, path string) string {
	return r.RepoRoot(repo) + "plain" + path
}

func (r *repoReverser) RepoTree(repo *models.Repo, path string, isDir bool) string {
	result := r.RepoRoot(repo) + "tree" + path
	hasSlash := result[len(result)-1] == '/'

	// Remove/add slash as needed
	if hasSlash && !isDir {
		result = result[:len(result)-1]
	} else if !hasSlash && isDir {
		result += "/"
	}

	return result
}

func (r *repoReverser) RepoLog(repo *models.Repo) string {
	return r.RepoRoot(repo) + "log/"
}

func (r *repoReverser) RepoRefs(repo *models.Repo) string {
	return r.RepoRoot(repo) + "refs/"
}

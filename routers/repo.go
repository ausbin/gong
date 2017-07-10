package routers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates"
	"code.austinjadams.com/gong/templates/url"
	url_ "net/url"
)

type Repo struct {
	cfg       *config.Global
	url       url.Reverser
	repo      models.Repo
	templates templates.Loader
}

func NewRepo(cfg *config.Global, url url.Reverser, repo models.Repo, templates templates.Loader) SubRouter {
	return &Repo{cfg, url, repo, templates}
}

func (r *Repo) ConfigureRouter(superRouter Router) {
	superRouter.Handle(r.url.RepoRoot(r.repo, ""), false,
		handlers.NewRepoRoot(r.cfg, r.url, r.repo, r.templates.Consumer("repo-root")))
	superRouter.Handle(r.url.RepoPlain(r.repo, "", "/"), true,
		handlers.NewRepoPlain(r.url, r.repo))
	superRouter.Handle(r.url.RepoTree(r.repo, "", "/", true), true,
		handlers.NewRepoTree(r.cfg, r.url, r.repo, r.templates.Consumer("repo-tree")))
	superRouter.Handle(r.url.RepoLog(r.repo, ""), false,
		handlers.NewRepoLog(r.cfg, r.url, r.repo, r.templates.Consumer("repo-log")))
	superRouter.Handle(r.url.RepoRefs(r.repo, ""), false,
		handlers.NewRepoRefs(r.cfg, r.url, r.repo, r.templates.Consumer("repo-refs")))
}

type repoReverser struct {
	repoPrefix string
}

func NewRepoReverser(repoPrefix string) url.RepoReverser {
	return &repoReverser{repoPrefix}
}

func (r *repoReverser) buildQueryString(repo models.Repo, branch string) string {
	if branch == "" {
		branch = repo.DefaultBranch()
	}

	if branch == repo.DefaultBranch() {
		return ""
	} else {
		return "?h=" + url_.QueryEscape(branch)
	}
}

func (r *repoReverser) repoRoot(repo models.Repo, branch string) string {
	return r.repoPrefix + "/" + repo.Name() + "/"
}

func (r *repoReverser) RepoRoot(repo models.Repo, branch string) string {
	return r.repoRoot(repo, branch) + r.buildQueryString(repo, branch)
}

func (r *repoReverser) repoPlain(repo models.Repo, branch string, path string) string {
	return r.repoRoot(repo, branch) + "plain" + path
}

func (r *repoReverser) RepoPlain(repo models.Repo, branch string, path string) string {
	return r.repoPlain(repo, branch, path) + r.buildQueryString(repo, branch)
}

func (r *repoReverser) repoTree(repo models.Repo, branch string, path string, isDir bool) string {
	result := r.repoRoot(repo, branch) + "tree" + path
	hasSlash := result[len(result)-1] == '/'

	// Remove/add slash as needed
	if hasSlash && !isDir {
		result = result[:len(result)-1]
	} else if !hasSlash && isDir {
		result += "/"
	}

	return result
}

func (r *repoReverser) RepoTree(repo models.Repo, branch string, path string, isDir bool) string {
	return r.repoTree(repo, branch, path, isDir) + r.buildQueryString(repo, branch)
}

func (r *repoReverser) repoLog(repo models.Repo, branch string) string {
	return r.repoRoot(repo, branch) + "log/"
}

func (r *repoReverser) RepoLog(repo models.Repo, branch string) string {
	return r.repoLog(repo, branch) + r.buildQueryString(repo, branch)
}

func (r *repoReverser) repoRefs(repo models.Repo, branch string) string {
	return r.repoRoot(repo, branch) + "refs/"
}

func (r *repoReverser) RepoRefs(repo models.Repo, branch string) string {
	return r.repoRefs(repo, branch) + r.buildQueryString(repo, branch)
}

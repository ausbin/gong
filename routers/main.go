package routers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/templates"
	"code.austinjadams.com/gong/templates/url"
	"net/http"
)

type main struct {
	*http.ServeMux
}

func NewMain(cfg config.Parser, templates templates.Loader) Router {
	m := &main{http.NewServeMux()}

	// Ideally, the downstream server (nginx, Apache, etc.) would handle
	// requests to /static/ instead, but this is useful for testing.
	m.Handle(cfg.Global().StaticPrefix+"/",
		http.StripPrefix(cfg.Global().StaticPrefix, http.FileServer(http.Dir(cfg.Global().StaticDir))))

	reverser := &reverser{
		RepoReverser: NewRepoReverser(cfg.Global().PathPrefix),
		root:         cfg.Global().PathPrefix,
		static:       cfg.Global().StaticPrefix,
	}

	// Register repository paths
	for _, repo := range cfg.Repos() {
		NewRepo(cfg.Global(), reverser, repo, templates).ConfigureRouter(m)
	}

	// Show a repository listing at /
	m.Handle(reverser.Root(),
		handlers.NewList(cfg.Global(), reverser, cfg.Repos(), templates.Get("list")))

	return m
}

type reverser struct {
	url.RepoReverser

	root, static string
}

func (r *reverser) Root() string {
	return r.root + "/"
}

func (r *reverser) Static(path string) string {
	return r.static + "/" + path
}

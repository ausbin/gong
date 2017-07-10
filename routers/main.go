package routers

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/templates"
	"code.austinjadams.com/gong/templates/url"
	"log"
	"net/http"
)

type main struct {
	mux *http.ServeMux
}

func NewMain(cfg config.Parser, templates templates.Loader) Router {
	m := &main{http.NewServeMux()}

	// Ideally, the downstream server (nginx, Apache, etc.) would handle
	// requests to /static/ instead, but this is useful for testing.
	m.HandleClassic(cfg.Global().StaticPrefix+"/",
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
	m.Handle(reverser.Root(), false,
		handlers.NewList(cfg.Global(), reverser, cfg.Repos(), templates.Consumer("list")))

	return m
}

func (m *main) HandleClassic(path string, handler http.Handler) {
	m.mux.Handle(path, handler)
}

func (m *main) Handle(path string, isSubtree bool, handler handlers.Handler) {
	m.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		// If this route is not a subtree (i.e., if it should not accept
		// sub-paths, like /X/Y/Z for a route /X/) and this request includes a
		// sub-path, then bail now and throw a 404.
		if !isSubtree && len(r.URL.Path) > len(path) {
			http.NotFound(w, r)
		} else {
			handler.Serve(newHandlerRequest(w, r, path))
		}
	})
}

func (m *main) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

type handlerRequest struct {
	writer      http.ResponseWriter
	request     *http.Request
	subtree     string
	written     bool
	querystring map[string]string
}

func newHandlerRequest(writer http.ResponseWriter, request *http.Request, registeredPath string) *handlerRequest {
	subtree := request.URL.Path[len(registeredPath)-1:]

	// Build the querystring map. For now, discard all but the last value
	// provided for a key
	querystring := make(map[string]string)
	for key, values := range request.URL.Query() {
		if len(values) > 0 {
			querystring[key] = values[len(values)-1]
		}
	}

	return &handlerRequest{writer, request, subtree, false, querystring}
}

func (hr *handlerRequest) Path() string                   { return hr.request.URL.Path }
func (hr *handlerRequest) Subtree() string                { return hr.subtree }
func (hr *handlerRequest) QueryString() map[string]string { return hr.querystring }
func (hr *handlerRequest) Write(data []byte) (int, error) {
	hr.written = true
	return hr.writer.Write(data)
}
func (hr *handlerRequest) Redirect(path string) {
	http.Redirect(hr.writer, hr.request, path, http.StatusMovedPermanently)
}
func (hr *handlerRequest) Error(err error) {
	log.Println(err)

	if !hr.written {
		http.Error(hr.writer, err.Error(), http.StatusInternalServerError)
	}
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

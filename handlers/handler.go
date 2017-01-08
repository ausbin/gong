package handlers

import (
	"net/http"
)

type Handler interface {
	Serve(w http.ResponseWriter, r *http.Request, info Info)
}

type Info interface {
	Subtree() string
}

type Func func(w http.ResponseWriter, r *http.Request, info Info)

func (f Func) Serve(w http.ResponseWriter, r *http.Request, info Info) {
	f(w, r, info)
}

func Adapter(handler http.Handler) Handler {
	return Func(func(w http.ResponseWriter, r *http.Request, _ Info) {
		handler.ServeHTTP(w, r)
	})
}

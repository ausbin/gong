package routers

import (
	"net/http"
)

type Router interface {
	Handle(path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type SubRouter interface {
	ConfigureRouter(superRouter Router)
}

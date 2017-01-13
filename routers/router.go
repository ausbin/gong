package routers

import (
	"code.austinjadams.com/gong/handlers"
	"net/http"
)

type Router interface {
	Handle(path string, isSubtree bool, handler handlers.Handler)
	HandleClassic(path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type SubRouter interface {
	ConfigureRouter(superRouter Router)
}

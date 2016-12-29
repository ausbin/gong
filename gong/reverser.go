package main

import (
	"code.austinjadams.com/gong/handlers"
)

type reverser struct {
	*handlers.RepoReverser

	root, static string
}

func (r *reverser) Root() string {
	return r.root
}

func (r *reverser) Static(path string) string {
	return r.static + path
}

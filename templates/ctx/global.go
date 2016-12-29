package ctx

import (
	"code.austinjadams.com/gong/templates/url"
)

// Global context
type Global interface {
	URL() url.Reverser
}

func NewGlobal(url url.Reverser) Global {
	return &global{url}
}

type global struct {
	url url.Reverser
}

func (g *global) URL() url.Reverser { return g.url }

package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/templates/url"
)

// Global context
type Global interface {
	Title() string
	Description() string
	URL() url.Reverser
	// Used in tests to compare contexts
	Equals(Global) bool
}

func NewGlobal(cfg *config.Global, url url.Reverser) Global {
	return &global{cfg, url}
}

type global struct {
	cfg *config.Global
	url url.Reverser
}

func (g *global) Title() string       { return g.cfg.Title }
func (g *global) Description() string { return g.cfg.Description }
func (g *global) URL() url.Reverser   { return g.url }
func (g *global) Equals(other Global) bool {
	return g.Title() == other.Title() &&
		g.Description() == other.Description() &&
		g.URL() == other.URL()
}

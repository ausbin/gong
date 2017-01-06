package ctx

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
)

type RepoTree interface {
	RepoGlobal

	IsRoot() bool
	Dirname() string
	SplitPath() []PathComponent
	Path() string
	IsListing() bool
	Files() []models.RepoFile
	Blob() string
}

func NewRepoTree(cfg *config.Global, url url.Reverser, repo *models.Repo, path string, isListing bool, files []models.RepoFile, blob string) RepoTree {
	return &repoTree{NewRepoGlobal(cfg, url, repo), path, NewPath(path), isListing, files, blob}
}

type repoTree struct {
	RepoGlobal

	path      string
	splitPath []PathComponent
	isListing bool
	files     []models.RepoFile
	blob      string
}

func (t *repoTree) Path() string             { return t.path }
func (t *repoTree) IsListing() bool          { return t.isListing }
func (t *repoTree) Files() []models.RepoFile { return t.files }
func (t *repoTree) Blob() string             { return t.blob }

func (t *repoTree) IsRoot() bool {
	return t.path == "/"
}

func (t *repoTree) SplitPath() []PathComponent {
	return t.splitPath
}

func (t *repoTree) Dirname() string {
	// If path is / or /X, return /
	if len(t.splitPath) <= 1 {
		return "/"
	} else {
		// Return the full path of the second-to-last path component. For
		// example, /W/X/Y/Z/ -> /W/X/Y/
		return t.splitPath[len(t.splitPath)-2].FullPath()
	}
}

type PathComponent interface {
	Name() string
	FullPath() string
	IsLast() bool
}

func NewPath(path string) []PathComponent {
	var components []PathComponent

	// Skip leading /
	nextIndex := 1

	for i := nextIndex; i <= len(path); i++ {
		// Don't create a component after a trailing slash
		if i == len(path) && nextIndex < len(path) || i < len(path) && path[i] == '/' {
			newComponent := &pathComponent{path, nextIndex, i - nextIndex}
			components = append(components, newComponent)

			nextIndex = i + 1
		}
	}

	return components
}

type pathComponent struct {
	full   string
	start  int
	length int
}

func (p *pathComponent) Name() string {
	return p.full[p.start : p.start+p.length]
}

func (p *pathComponent) FullPath() string {
	return p.full[:p.start+p.length]
}

func (p *pathComponent) IsLast() bool {
	return p.start+p.length >= len(p.full)-1
}

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
	IsBinary() bool
	IsImage() bool
	Files() []models.RepoFile
	Blob() string
}

func NewRepoTree(cfg *config.Global, url url.Reverser, repo models.Repo, branch string, branches []string, path string, isListing bool, isBinary bool, isImage bool, files []models.RepoFile, blob string) RepoTree {
	return &repoTree{NewRepoGlobal(cfg, url, repo, branch, branches), path, NewPath(path), isListing, isBinary, isImage, files, blob}
}

type repoTree struct {
	RepoGlobal

	path      string
	splitPath []PathComponent
	isListing bool
	isBinary  bool
	isImage   bool
	files     []models.RepoFile
	blob      string
}

func (t *repoTree) Path() string             { return t.path }
func (t *repoTree) IsListing() bool          { return t.isListing }
func (t *repoTree) IsBinary() bool           { return t.isBinary }
func (t *repoTree) IsImage() bool            { return t.isImage }
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

// XXX Don't
func (t *repoTree) pathComponentsEqual(a, b []PathComponent) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

func (t *repoTree) filesEqual(a, b []models.RepoFile) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

func (t *repoTree) Equals(other Global) bool {
	otherTree, ok := other.(RepoRoot)

	return ok && t.RepoGlobal.Equals(other) &&
		t.IsRoot() == otherTree.IsRoot() &&
		t.Dirname() == otherTree.Dirname() &&
		t.pathComponentsEqual(t.SplitPath(), otherTree.SplitPath()) &&
		t.Path() == otherTree.Path() &&
		t.IsListing() == otherTree.IsListing() &&
		t.IsBinary() == otherTree.IsBinary() &&
		t.IsImage() == otherTree.IsImage() &&
		t.filesEqual(t.Files(), otherTree.Files()) &&
		t.Blob() == otherTree.Blob()
}

type PathComponent interface {
	Name() string
	FullPath() string
	IsLast() bool
	Equals(PathComponent) bool
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

func (p *pathComponent) Equals(other PathComponent) bool {
	return p.Name() == other.Name() &&
		p.FullPath() == other.FullPath() &&
		p.IsLast() == other.IsLast()
}

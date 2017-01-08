package templates

// Add new templates here
// Order: name, parent name
var templatesToLoad = []loaderTempl{
	{"layout", ""},
	{"list", "layout"},
	{"repo", "layout"},
	{"repo-tree", "repo"},
	{"repo-root", "repo-tree"},
	{"repo-refs", "repo"},
	{"repo-log", "repo"},
}

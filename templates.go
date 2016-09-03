package gong

import (
	"html/template"
)

func LoadTemplates(dir string) (map[string]*template.Template, error) {
	m := make(map[string]*template.Template)

	layout, err := template.ParseFiles(dir + "/layout.html")
	if err != nil {
		return nil, err
	}

	// The current implementation of Clone() does not return an error
	// under any circumstances, so keep only the cloned template
	clone, _ := layout.Clone()
	repo, err := clone.ParseFiles(dir + "/repo.html")
	if err != nil {
		return nil, err
	}

	clone, _ = repo.Clone()
	m["repo-tree"], err = clone.ParseFiles(dir + "/repo-tree.html")
	if err != nil {
		return nil, err
	}

	return m, nil
}

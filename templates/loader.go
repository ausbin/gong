package templates

import (
	"errors"
	"html/template"
)

type Loader interface {
	Get(name string) *template.Template
}

func NewLoader(dir string) (Loader, error) {
	return newLoader(dir, templatesToLoad)
}

func newLoader(dir string, templates []loaderTempl) (*loader, error) {
	l := &loader{dir, make(map[string]*template.Template)}
	err := l.loadAll(templates)

	// Return a nil Loader when an error occurs
	if err != nil {
		l = nil
	}

	return l, err
}

type loader struct {
	dir       string
	templates map[string]*template.Template
}

func (l *loader) loadAll(all []loaderTempl) error {
	for _, t := range all {
		if err := t.loadInto(l); err != nil {
			return err
		}
	}

	return nil
}

func (l *loader) has(name string) bool {
	_, exists := l.templates[name]
	return exists
}

func (l *loader) get(name string) (*template.Template, bool) {
	t, ok := l.templates[name]
	return t, ok
}

func (l *loader) add(name string, templ *template.Template) {
	l.templates[name] = templ
}

// Same as get(), but panics if the template does not exist
func (l *loader) Get(name string) *template.Template {
	t, ok := l.get(name)

	if !ok {
		panic("template " + name + " does not exist")
	} else {
		return t
	}
}

type loaderTempl struct {
	name   string
	parent string
}

func (t *loaderTempl) loadInto(l *loader) (err error) {
	if l.has(t.name) {
		return errors.New("template " + t.name + " already exists")
	}

	var templ *template.Template
	path := l.dir + "/" + t.name + ".html"

	if t.parent == "" {
		// This is the root template
		templ, err = template.ParseFiles(path)
	} else {
		// This is a child template
		parentTempl, ok := l.get(t.parent)

		if ok {
			templ, err = template.Must(parentTempl.Clone()).ParseFiles(path)
		} else {
			err = errors.New("parent template " + t.parent + " does not exist")
		}
	}

	if err != nil {
		return
	}

	l.add(t.name, templ)

	return
}

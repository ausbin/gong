package templates

import (
	"bytes"
	"testing"
)

const (
	templateDir   = "tests"
	templateBogus = "bogusbogusbogus"
	templateEmpty = "empty"
	templateFoo   = "foo"
	templateBar   = "bar"
)

func newBogusLoader(templates []loaderTempl) (*loader, error) {
	return newLoader(templateDir, templates)
}

func TestEmptyLoaderGet(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Get()ing from an unitialized Loader did not cause a panic")
		}
	}()

	loader, err := newBogusLoader(nil)

	if err != nil {
		t.Error("creating an empty loader failed")
	}

	loader.Get(templateEmpty)
}

func TestEmptyLoaderConsumer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("calling Consumer()ing on an unitialized Loader did not cause a panic")
		}
	}()

	loader, err := newBogusLoader(nil)

	if err != nil {
		t.Error("creating an empty loader failed")
	}

	loader.Get(templateEmpty)
}

func TestLoadBogusTemplate(t *testing.T) {
	_, err := newBogusLoader([]loaderTempl{{templateBogus, ""}})

	if err == nil {
		t.Error("loading a bogus template succeeded")
	}
}

func TestBogusParentTemplate(t *testing.T) {
	_, err := newBogusLoader([]loaderTempl{{templateEmpty, templateBogus}})

	if err == nil {
		t.Error("loading a template with a bogus parent succeded")
	}
}

func TestLoaderHas(t *testing.T) {
	loader, err := newBogusLoader([]loaderTempl{
		{templateEmpty, ""},
		{templateFoo, ""},
		{templateBar, templateFoo},
	})

	if err != nil {
		t.Error("loading valid templates failed. error:", err)
	}

	if loader.has(templateBogus) {
		t.Error("loader incorrectly claims to have ", templateBogus)
	}
	if !loader.has(templateEmpty) {
		t.Error("loader incorrectly claims not to have ", templateEmpty)
	}
	if !loader.has(templateFoo) {
		t.Error("loader incorrectly claims not to have ", templateFoo)
	}
	if !loader.has(templateBar) {
		t.Error("loader incorrectly claims not to have ", templateBar)
	}
}

func TestEmptyTemplate(t *testing.T) {
	loader, err := newBogusLoader([]loaderTempl{{templateEmpty, ""}})

	if err != nil {
		t.Error("loading a valid template failed. error:", err)
	}

	buf := new(bytes.Buffer)
	templ := loader.Get(templateEmpty)
	err = templ.Execute(buf, nil)

	if err != nil {
		t.Error("executing an empty template failed")
	}

	if buf.String() != "" {
		t.Error("executing an empty template yielded a nonempty string")
	}
}

func TestSimpleTemplate(t *testing.T) {
	loader, err := newBogusLoader([]loaderTempl{
		{templateEmpty, ""},
		{templateFoo, ""},
	})

	if err != nil {
		t.Error("loading valid templates failed. error:", err)
	}

	buf := new(bytes.Buffer)
	templ := loader.Get(templateFoo)
	err = templ.Execute(buf, struct{ Foo string }{"blah<><><>blah"})

	if err != nil {
		t.Error("executing an empty template failed")
	}

	if want, got := "Foo is blah&lt;&gt;&lt;&gt;&lt;&gt;blah\n\n", buf.String(); want != got {
		t.Error("executing a simple template yielded wrong output. expected", want, "got", got)
	}
}

func TestChildTemplate(t *testing.T) {
	loader, err := newBogusLoader([]loaderTempl{
		{templateEmpty, ""},
		{templateFoo, ""},
		{templateBar, templateFoo},
	})

	if err != nil {
		t.Error("loading valid templates failed. error:", err)
	}

	buf := new(bytes.Buffer)
	templ := loader.Get(templateBar)
	err = templ.Execute(buf, struct{ Foo string }{"kaboom"})

	if err != nil {
		t.Error("executing an empty template failed")
	}

	if want, got := "Foo is kaboom\nhello hello hello\n", buf.String(); want != got {
		t.Error("executing a child template yielded wrong output. expected", want, "got", got)
	}
}

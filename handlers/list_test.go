package handlers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/ctx"
	"testing"
)

func TestEmptyList(t *testing.T) {
	global := newBogusGlobal()
	reverser := newBogusReverser()

	ctxCollector := newBogusConsumer()
	NewList(global, reverser, nil, ctxCollector).Serve(newBogusRequest())

	want := ctx.NewList(global, reverser, nil)
	got := ctxCollector.ctx

	if got == nil {
		t.Error("ctx is nil")
	} else if !got.Equals(want) {
		t.Error("ctx", got, "differs from expected", want)
	}
}

func TestBasicList(t *testing.T) {
	global := newBogusGlobal()
	reverser := newBogusReverser()
	repos := []models.Repo{
		models.NewRepo("gong", "a git repository viewer", "/x/y/z", "master"),
		models.NewRepo("foo", "bar", "y/z", "dev"),
	}

	ctxCollector := newBogusConsumer()
	NewList(global, reverser, repos, ctxCollector).Serve(newBogusRequest())

	want := ctx.NewList(global, reverser, repos)
	got := ctxCollector.ctx

	if got == nil {
		t.Error("ctx is nil")
	} else if !got.Equals(want) {
		t.Error("ctx", got, "differs from expected", want)
	}
}

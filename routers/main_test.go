package routers

import (
	"testing"
)

func makeReverser() *reverser {
	return &reverser{
		RepoReverser: nil,
		root:         "/x/yyy/z",
		static:       "/a/b/ccc",
	}
}

func TestReverserStatic(t *testing.T) {
	rev := makeReverser()

	if rev.Static("bob.css") != "/a/b/ccc/bob.css" {
		t.Error("wrong path for static file bob.css")
	}

	if rev.Static("foo/bar") != "/a/b/ccc/foo/bar" {
		t.Error("wrong path for static file foo/bar")
	}

	if rev.Static("/heyyy") != "/a/b/ccc//heyyy" {
		t.Error("wrong path for static file /heyyy. should be static_prefix//heyyy")
	}

	// Now try again, except with an empty prefix
	rev.static = ""

	if rev.Static("larry.css") != "/larry.css" {
		t.Error("for an empty static prefix, Static(\"X\") should return \"/X\"")
	}

	if rev.Static("") != "/" {
		t.Error("for an empty static prefix, Static(\"\") should return \"/\"")
	}

	if rev.Static("/ha/ha/ha/ha/ho") != "//ha/ha/ha/ha/ho" {
		t.Error("for an empty static prefix, Static(\"/X\") should return \"//X\"")
	}
}

func TestReverserRoot(t *testing.T) {
	rev := makeReverser()

	if rev.Root() != "/x/yyy/z/" {
		t.Error("Root() does not return trailing slash. for some prefix /X, Root() should return /X/")
	}

	// Now try again, except with an empty prefix
	rev.root = ""

	if rev.Root() != "/" {
		t.Error("for an empty prefix, Root() should return /")
	}
}

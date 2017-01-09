package routers

import (
	"code.austinjadams.com/gong/models"
	"code.austinjadams.com/gong/templates/url"
	"testing"
)

func makeRepoReverser1() url.RepoReverser {
	return NewRepoReverser("/foo/bar/baz")
}

func makeRepoReverser2() url.RepoReverser {
	return NewRepoReverser("")
}

func makeRepo1() *models.Repo {
	// Can't call models.NewRepo() since we don't need/want to open an actual
	// repository with git2go
	return &models.Repo{
		Name: "linux",
	}
}

func makeRepo2() *models.Repo {
	// Can't call models.NewRepo() since we don't need/want to open an actual
	// repository with git2go
	return &models.Repo{
		Name: "gongitymcgongicus",
	}
}

func TestRepoReverserRoot(t *testing.T) {
	rev1 := makeRepoReverser1()
	rev2 := makeRepoReverser2()
	repo1 := makeRepo1()
	repo2 := makeRepo2()

	if rev1.RepoRoot(repo1) != "/foo/bar/baz/linux/" {
		t.Error("RepoRoot() returned wrong value")
	}

	if rev1.RepoRoot(repo2) != "/foo/bar/baz/gongitymcgongicus/" {
		t.Error("RepoRoot() returned wrong value")
	}

	if rev2.RepoRoot(repo1) != "/linux/" {
		t.Error("RepoRoot() returned wrong value")
	}

	if rev2.RepoRoot(repo2) != "/gongitymcgongicus/" {
		t.Error("RepoRoot() returned wrong value")
	}
}

func TestRepoReverserPlain(t *testing.T) {
	rev1 := makeRepoReverser1()
	rev2 := makeRepoReverser2()
	repo1 := makeRepo1()
	repo2 := makeRepo2()

	tests := []struct {
		rev url.RepoReverser
		repo *models.Repo
		path string
		expected string
	}{
		{rev1, repo1, "/x", "/foo/bar/baz/linux/plain/x"},
		{rev1, repo1, "/bananas/strawberry.txt", "/foo/bar/baz/linux/plain/bananas/strawberry.txt"},
		{rev1, repo2, "/x", "/foo/bar/baz/gongitymcgongicus/plain/x"},
		{rev1, repo2, "/bananas/strawberry.txt", "/foo/bar/baz/gongitymcgongicus/plain/bananas/strawberry.txt"},
		{rev2, repo1, "/x", "/linux/plain/x"},
		{rev2, repo1, "/bananas/strawberry.txt", "/linux/plain/bananas/strawberry.txt"},
		{rev2, repo2, "/x", "/gongitymcgongicus/plain/x"},
		{rev2, repo2, "/bananas/strawberry.txt", "/gongitymcgongicus/plain/bananas/strawberry.txt"},
	}

	for _, test := range tests {
		if got := test.rev.RepoPlain(test.repo, test.path); got != test.expected {
			t.Error("RepoPlain() returned wrong value",
			        "RepoPlain(", test.repo, test.path, ") →", got, "!=", test.expected)
		}
	}
}

func TestRepoReverserTree(t *testing.T) {
	rev1 := makeRepoReverser1()
	rev2 := makeRepoReverser2()
	repo1 := makeRepo1()
	repo2 := makeRepo2()

	tests := []struct {
		rev url.RepoReverser
		repo *models.Repo
		isDir bool
		path string
		expected string
	}{
		// Nonempty prefix
		// Not a dir, no slash
		{rev1, repo1, false, "/x", "/foo/bar/baz/linux/tree/x"},
		{rev1, repo1, false, "/bananas/strawberry.txt", "/foo/bar/baz/linux/tree/bananas/strawberry.txt"},
		{rev1, repo2, false, "/x", "/foo/bar/baz/gongitymcgongicus/tree/x"},
		{rev1, repo2, false, "/bananas/strawberry.txt", "/foo/bar/baz/gongitymcgongicus/tree/bananas/strawberry.txt"},
		// Not a dir, slash
		{rev1, repo1, false, "/x/", "/foo/bar/baz/linux/tree/x"},
		{rev1, repo1, false, "/bananas/strawberry.txt/", "/foo/bar/baz/linux/tree/bananas/strawberry.txt"},
		{rev1, repo2, false, "/x/", "/foo/bar/baz/gongitymcgongicus/tree/x"},
		{rev1, repo2, false, "/bananas/strawberry.txt/", "/foo/bar/baz/gongitymcgongicus/tree/bananas/strawberry.txt"},
		// Dir, no slash
		{rev1, repo1, true, "/x", "/foo/bar/baz/linux/tree/x/"},
		{rev1, repo1, true, "/bananas/strawberry.txt", "/foo/bar/baz/linux/tree/bananas/strawberry.txt/"},
		{rev1, repo2, true, "/x", "/foo/bar/baz/gongitymcgongicus/tree/x/"},
		{rev1, repo2, true, "/bananas/strawberry.txt", "/foo/bar/baz/gongitymcgongicus/tree/bananas/strawberry.txt/"},
		// Dir, slash
		{rev1, repo1, true, "/x/", "/foo/bar/baz/linux/tree/x/"},
		{rev1, repo1, true, "/bananas/strawberry.txt/", "/foo/bar/baz/linux/tree/bananas/strawberry.txt/"},
		{rev1, repo2, true, "/x/", "/foo/bar/baz/gongitymcgongicus/tree/x/"},
		{rev1, repo2, true, "/bananas/strawberry.txt/", "/foo/bar/baz/gongitymcgongicus/tree/bananas/strawberry.txt/"},
		// Empty prefix
		// Not a dir, no slash
		{rev2, repo1, false, "/x", "/linux/tree/x"},
		{rev2, repo1, false, "/bananas/strawberry.txt", "/linux/tree/bananas/strawberry.txt"},
		{rev2, repo2, false, "/x", "/gongitymcgongicus/tree/x"},
		{rev2, repo2, false, "/bananas/strawberry.txt", "/gongitymcgongicus/tree/bananas/strawberry.txt"},
		// Not a dir, slash
		{rev2, repo1, false, "/x/", "/linux/tree/x"},
		{rev2, repo1, false, "/bananas/strawberry.txt/", "/linux/tree/bananas/strawberry.txt"},
		{rev2, repo2, false, "/x/", "/gongitymcgongicus/tree/x"},
		{rev2, repo2, false, "/bananas/strawberry.txt/", "/gongitymcgongicus/tree/bananas/strawberry.txt"},
		// Dir, no slash
		{rev2, repo1, true, "/x", "/linux/tree/x/"},
		{rev2, repo1, true, "/bananas/strawberry.txt", "/linux/tree/bananas/strawberry.txt/"},
		{rev2, repo2, true, "/x", "/gongitymcgongicus/tree/x/"},
		{rev2, repo2, true, "/bananas/strawberry.txt", "/gongitymcgongicus/tree/bananas/strawberry.txt/"},
		// Dir, slash
		{rev2, repo1, true, "/x/", "/linux/tree/x/"},
		{rev2, repo1, true, "/bananas/strawberry.txt/", "/linux/tree/bananas/strawberry.txt/"},
		{rev2, repo2, true, "/x/", "/gongitymcgongicus/tree/x/"},
		{rev2, repo2, true, "/bananas/strawberry.txt/", "/gongitymcgongicus/tree/bananas/strawberry.txt/"},
	}

	for _, test := range tests {
		if got := test.rev.RepoTree(test.repo, test.path, test.isDir); got != test.expected {
			t.Error("RepoTree() returned wrong value.",
			        "RepoTree(", test.repo, test.path, test.isDir, ") →", got, "!=", test.expected)
		}
	}
}

func TestRepoReverserLog(t *testing.T) {
	rev1 := makeRepoReverser1()
	rev2 := makeRepoReverser2()
	repo1 := makeRepo1()
	repo2 := makeRepo2()

	if rev1.RepoLog(repo1) != "/foo/bar/baz/linux/log/" {
		t.Error("RepoLog() returned wrong value")
	}

	if rev1.RepoLog(repo2) != "/foo/bar/baz/gongitymcgongicus/log/" {
		t.Error("RepoLog() returned wrong value")
	}

	if rev2.RepoLog(repo1) != "/linux/log/" {
		t.Error("RepoLog() returned wrong value")
	}

	if rev2.RepoLog(repo2) != "/gongitymcgongicus/log/" {
		t.Error("RepoLog() returned wrong value")
	}
}

func TestRepoReverserRefs(t *testing.T) {
	rev1 := makeRepoReverser1()
	rev2 := makeRepoReverser2()
	repo1 := makeRepo1()
	repo2 := makeRepo2()

	if rev1.RepoRefs(repo1) != "/foo/bar/baz/linux/refs/" {
		t.Error("RepoRefs() returned wrong value")
	}

	if rev1.RepoRefs(repo2) != "/foo/bar/baz/gongitymcgongicus/refs/" {
		t.Error("RepoRefs() returned wrong value")
	}

	if rev2.RepoRefs(repo1) != "/linux/refs/" {
		t.Error("RepoRefs() returned wrong value")
	}

	if rev2.RepoRefs(repo2) != "/gongitymcgongicus/refs/" {
		t.Error("RepoRefs() returned wrong value")
	}
}

package config

import (
	"testing"
)

func assertStr(t *testing.T, desc, want, got string) {
	if want != got {
		t.Error(desc, "is wrong. Expected", want, "but got", got)
	}
}

func assertUint(t *testing.T, desc string, want, got uint) {
	if want != got {
		t.Error(desc, "is wrong. Expected", want, "but got", got)
	}
}

func TestGlobalDefaults(t *testing.T) {
	cfg, err := NewBogusParser([]byte(""))

	if err != nil {
		t.Log("creating empty config parser failed. error:", err)
		t.FailNow()
	}

	global := cfg.Global()

	assertStr(t, "Default Global().Addr", "127.0.0.1", global.Addr)
	assertUint(t, "Default Global().Port", 8050, global.Port)
	assertStr(t, "Default Global().Title", "gong", global.Title)
	assertStr(t, "Default Global().Description", "a gong instance", global.Description)
	assertStr(t, "Default Global().PathPrefix", "", global.PathPrefix)
	assertStr(t, "Default Global().StaticPrefix", "/static", global.StaticPrefix)
	assertStr(t, "Default Global().StaticDir", "static", global.StaticDir)
	assertStr(t, "Default Global().TemplateDir", "templates", global.TemplateDir)
}

func TestGlobalExample(t *testing.T) {
	cfg, err := NewBogusParser([]byte(`
		title = A GOOD WEB SITE !
		description = ok it's mediocre
		addr = 8.8.8.8
		port = 8001
		path_prefix = /code/
		static_prefix = /code/staticFiles/
		static_dir = /var/xyz/gong/static
		template_dir = /var/xyz/gong/templates
	`))

	if err != nil {
		t.Log("creating a config parser against valid config failed. error:", err)
		t.FailNow()
	}

	global := cfg.Global()

	assertStr(t, "Global().Addr", "8.8.8.8", global.Addr)
	assertUint(t, "Global().Port", 8001, global.Port)
	assertStr(t, "Global().Title", "A GOOD WEB SITE !", global.Title)
	assertStr(t, "Global().Description", "ok it's mediocre", global.Description)
	assertStr(t, "Global().PathPrefix", "/code", global.PathPrefix)
	assertStr(t, "Global().StaticPrefix", "/code/staticFiles", global.StaticPrefix)
	assertStr(t, "Global().StaticDir", "/var/xyz/gong/static", global.StaticDir)
	assertStr(t, "Global().TemplateDir", "/var/xyz/gong/templates", global.TemplateDir)
}

func TestInvalidPort(t *testing.T) {
	_, err := NewBogusParser([]byte(`
		title = a dope git webzone
		port = this is invalid
		addr = 8.8.8.8
	`))

	if err == nil {
		t.Error("creating a config parser against config with invalid port number succeeded")
	}
}

func TestNoRepos(t *testing.T) {
	cfg, err := NewBogusParser([]byte(""))

	if err != nil {
		t.Log("creating empty config parser failed. error:", err)
		t.FailNow()
	}

	if len(cfg.Repos()) > 0 {
		t.Error("empty config file includes more than zero repositories")
	}
}

func TestReposExample(t *testing.T) {
	repos := map[string]struct {
		description, path, defbranch string
	}{
		"foobity":      {"foobities the bar", "/var/git/foobity", "dev"},
		"a handy tool": {"", "/opt/handy", "dev2"},
		"🔥":            {"it's lit", "lit", "master"},
		"a-really-long-repository-name-wew-buddy": {"", "/a/painfully/long/repository/path/oh/yeah/keep/going/", "master"},
	}

	cfg, err := NewBogusParser([]byte(`
		[foobity]
		description = foobities the bar
		path = /var/git/foobity
		defbranch = dev

		[a handy tool]
		path = /opt/handy
		defbranch = dev2

		[🔥]
		path = lit
		description = it's lit

		[a-really-long-repository-name-wew-buddy]
		path = /a/painfully/long/repository/path/oh/yeah/keep/going/
	`))

	if err != nil {
		t.Log("creating a config parser against valid config failed. error:", err)
		t.FailNow()
	}

	for _, parsed := range cfg.Repos() {
		wanted, ok := repos[parsed.Name()]

		if !ok {
			t.Error("Unexpected repository", parsed.Name())
		} else {
			assertStr(t, "Repo description", wanted.description, parsed.Description())
			assertStr(t, "Repo path", wanted.path, parsed.Path())
			assertStr(t, "Repo default branch", wanted.defbranch, parsed.DefaultBranch())
		}
	}
}

func TestRepoWithoutPath(t *testing.T) {
	_, err := NewBogusParser([]byte(`
		title = this should fail
		description = since repos need a path

		[broken]
	`))

	if err == nil {
		t.Error("a repo section without a path succeeded")
	}
}

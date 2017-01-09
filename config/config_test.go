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
	cfg, err := NewParserBytes([]byte(""))

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
	cfg, err := NewParserBytes([]byte(`
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

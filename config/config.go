package config

import (
	"code.austinjadams.com/gong/models"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

type Parser interface {
	Global() *Global
	Repos() []*models.Repo
}

func NewParser(path string) (Parser, error) {
	p := &parser{}
	err := p.parse(path)

	if err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

type Global struct {
	// Bind address+port
	Addr string
	Port uint
	// Server configuration
	PathPrefix   string
	StaticPrefix string
	StaticDir    string
	TemplateDir  string
}

// Return address and port as an ADDR:PORT pair
func (g *Global) BindInfo() string {
	return g.Addr + ":" + strconv.FormatUint(uint64(g.Port), 10)
}

type parser struct {
	cfg    *ini.File
	global *Global
	repos  []*models.Repo
}

func (p *parser) parse(path string) error {
	var err error
	p.cfg, err = ini.Load(path)

	if err != nil {
		return err
	}

	if err = p.parseGlobal(); err != nil {
		return err
	}

	if err = p.parseRepos(); err != nil {
		return err
	}

	return nil
}

func (p *parser) parseGlobal() error {
	// Default section (which will always exist) contains global settings
	section, _ := p.cfg.GetSection("")

	p.global = &Global{
		Addr:         section.Key("addr").MustString("127.0.0.1"),
		Port:         section.Key("port").MustUint(8050),
		PathPrefix:   strings.TrimSuffix(section.Key("path_prefix").MustString("/"), "/"),
		StaticPrefix: strings.TrimSuffix(section.Key("static_prefix").MustString("/static"), "/"),
		StaticDir:    section.Key("static_dir").MustString("static"),
		TemplateDir:  section.Key("template_dir").MustString("templates"),
	}

	return nil
}

func (p *parser) parseRepos() error {
	sections := p.cfg.Sections()

	for _, section := range sections {
		// Ignore the default section, which holds global settings
		if section.Name() == ini.DEFAULT_SECTION {
			continue
		}

		newRepo, err := models.NewRepo(
			section.Name(),
			section.Key("description").MustString(""),
			section.Key("path").String(),
		)

		if err != nil {
			return err
		}

		p.repos = append(p.repos, newRepo)
	}

	return nil
}

func (p *parser) Global() *Global {
	return p.global
}

func (p *parser) Repos() []*models.Repo {
	return p.repos
}

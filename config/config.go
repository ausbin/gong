package config

import (
	"code.austinjadams.com/gong/models"
	"errors"
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
	err := p.readFile(path)

	if err != nil {
		return nil, err
	}

	err = p.parse(true)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// Parse config from a byte slice, without attempting to open repositories
// paths. Used in unit tests.
func NewBogusParser(data []byte) (Parser, error) {
	p := &parser{}
	err := p.readBytes(data)

	if err != nil {
		return nil, err
	}

	err = p.parse(false)

	if err != nil {
		return nil, err
	}

	return p, nil
}

type Global struct {
	// Bind address+port
	Addr string
	Port uint
	// Server configuration
	Title        string
	Description  string
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

func (p *parser) readBytes(data []byte) error {
	var err error
	p.cfg, err = ini.Load(data)
	return err
}

func (p *parser) readFile(path string) error {
	var err error
	p.cfg, err = ini.Load(path)
	return err
}

func (p *parser) parse(openRepos bool) error {
	err := p.parseGlobal()

	if err != nil {
		return err
	}

	err = p.parseRepos(openRepos)

	if err != nil {
		return err
	}

	return nil
}

func (p *parser) parseGlobal() error {
	// Default section (which will always exist) contains global settings
	section, _ := p.cfg.GetSection("")

	var port uint

	// Handle three cases: (1) no port given, so use default; (2) valid port
	// given; and (3) invalid port given, so return the error.
	// go-ini's MustUint() fails painfully at handling these, particularly (3).
	// Indeed, if a user supplies an invalid port, it silently fails, falling
	// back to the default I supply here.
	if val := section.Key("port").String(); val == "" {
		port = 8050
	} else if port_, err := strconv.ParseUint(val, 10, 0); err == nil {
		port = uint(port_)
	} else {
		return err
	}

	p.global = &Global{
		Addr:         section.Key("addr").MustString("127.0.0.1"),
		Port:         port,
		Title:        section.Key("title").MustString("gong"),
		Description:  section.Key("description").MustString("a gong instance"),
		PathPrefix:   strings.TrimSuffix(section.Key("path_prefix").MustString("/"), "/"),
		StaticPrefix: strings.TrimSuffix(section.Key("static_prefix").MustString("/static"), "/"),
		StaticDir:    section.Key("static_dir").MustString("static"),
		TemplateDir:  section.Key("template_dir").MustString("templates"),
	}

	return nil
}

func (p *parser) parseRepos(openRepos bool) error {
	sections := p.cfg.Sections()

	for _, section := range sections {
		// Ignore the default section, which holds global settings
		if section.Name() == ini.DEFAULT_SECTION {
			continue
		}

		name := section.Name()
		path := section.Key("path").String()

		if path == "" {
			return errors.New("Must specify path for repository " + name)
		}

		newRepo := models.NewRepo(
			name,
			section.Key("description").MustString(""),
			path,
			section.Key("defbranch").MustString("master"),
		)

		if openRepos {
			err := newRepo.Open()

			if err != nil {
				return err
			}
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

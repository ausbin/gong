package main

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/handlers"
	"code.austinjadams.com/gong/templates"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args)-1 != 1 {
		log.Fatalln("you must pass only the path to a configuration file")
	}

	cfg, err := config.NewParser(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	templates, err := templates.NewLoader(cfg.Global().TemplateDir)

	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	// Ideally, the downstream server (nginx, Apache, etc.) would handle
	// requests to /static/ instead, but this is useful for testing.
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	// Register repository paths
	for _, repo := range cfg.Repos() {
		path := "/" + repo.Name + "/tree"
		mux.Handle(path+"/", http.StripPrefix(path, handlers.NewTree(repo, templates.Get("repo-tree"))))
	}

	bindTo := cfg.Global().Addr + ":" + strconv.FormatUint(uint64(cfg.Global().Port), 10)
	err = http.ListenAndServe(bindTo, mux)

	if err != nil {
		log.Fatalln(err)
	}
}

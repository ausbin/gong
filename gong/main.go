package main

import (
	"code.austinjadams.com/gong/config"
	"code.austinjadams.com/gong/routers"
	"code.austinjadams.com/gong/templates"
	"log"
	"net/http"
	"os"
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

	router := routers.NewMain(cfg, templates)
	err = http.ListenAndServe(cfg.Global().BindInfo(), router)

	if err != nil {
		log.Fatalln(err)
	}
}

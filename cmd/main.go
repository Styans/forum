package main

import (
	"flag"
	"forum/configs"
	"forum/internal/app"
	"forum/internal/handlers"
	"forum/internal/render"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/pkg/client/sqlite"
	"log"
)

func main() {
	log.Println("wait a minute...")

	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	cfg, err := configs.GetConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := sqlite.OpenDB(cfg.DB.DSN)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	template, err := render.NewTemplateHTML(cfg.TemplateDir)
	if err != nil {
		log.Println(err)
		return
	}

	handler := handlers.NewHandler(service, template)

	err = app.Server(cfg, handler.Routes())

	if err != nil {

		log.Println("Ooopss...\n", err)
		return
	}
}

package main

import (
	"flag"
	"fmt"
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
	configPath := flag.String("config", "config.json", "path to config file")

	flag.Parse()

	cfg, err := configs.GetConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(cfg.DB.DSN)
	db, err := sqlite.OpenDB(cfg.DB.DSN)
	if err != nil {
		log.Fatal("hello")
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
		log.Println(err)
	}
}

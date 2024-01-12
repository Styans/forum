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
	configPath := flag.String("config", "config.json", "path to config file")

	flag.Parse()

	cfg, err := configs.GetConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}
	// db, err := sqlite3.OpenDB(cfg.DB.DSN)
	db, err := sqlite.OpenDB(cfg.DB.DSN)

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	template := render.NewTemplateHTML()

	handler := handlers.NewHandler(service, template)

	err = app.Server(cfg, handler.Routes())

	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"flag"
	"forun/configs"
	"forun/internal/app"
	"forun/internal/handlers"
	"forun/internal/render"
	"forun/internal/service"
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
	// service := service.NewService()

	// template := render.NewTemplateHTML()

	handler := handlers.NewHandler(
		service.NewService(),
		render.NewTemplateHTML(),
	)

	err = app.Server(cfg, handler.Routes())

	if err != nil {
		log.Println(err)
	}
}

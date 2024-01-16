package app

import (
	"forum/configs"
	"log"
	"net/http"
	"time"
)

func Server(cfg *configs.Config, handler http.Handler) error {
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("The server is running, you can use it using this link http://localhost:8080/")
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

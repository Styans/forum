package app

import (
	"forun/configs"
	"net/http"
)

func Server(cfg *configs.Config, handler http.Handler) error {
	
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

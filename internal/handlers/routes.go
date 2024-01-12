package handlers

import "net/http"

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	//add a css file to route
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}

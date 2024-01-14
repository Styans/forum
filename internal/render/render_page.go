package render

import (
	"log"
	"net/http"
)

func (t *TemplatesHTML) Render(w http.ResponseWriter, r *http.Request, name string, data *PageData) {
	tmlp, ok := (*t)[name]
	if !ok {
		log.Println("#")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmlp.Execute(w, data)
}

package render

import (
	"bytes"
	"fmt"
	"net/http"
)

func (t *TemplatesHTML) Render(w http.ResponseWriter, r *http.Request, name string, data *PageData) {
	tmlp, ok := (*t)[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	err := tmlp.Execute(buf, data)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

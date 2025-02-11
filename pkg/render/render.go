package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.New(tmpl).ParseFiles("web/templates/" + tmpl)
	if err != nil {
		fmt.Println("Error rendering:", err)
		return err
	}

	return t.Execute(w, data)
}

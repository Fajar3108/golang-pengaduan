package helpers

import (
	"net/http"
	"text/template"
)

// Render HTML File with partials file
func View(rw http.ResponseWriter, viewName string, data map[string]interface{}, files ...string) {
	tmpl := template.Must(template.ParseFiles(files...))

	err := tmpl.ExecuteTemplate(rw, viewName, data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

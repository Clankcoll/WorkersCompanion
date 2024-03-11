package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to your index.html file
	tmplPath, err := filepath.Abs("../../web/static/index.html")
	if err != nil {
		http.Error(w, "Failed to find the index.html file", http.StatusInternalServerError)
		return
	}

	// Parse the index.html file
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Failed to load index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template without passing any data
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to render index.html: "+err.Error(), http.StatusInternalServerError)
	}
}

package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

// DisplayForm handles GET requests to display the form
func WorkTimeCalculator(w http.ResponseWriter, r *http.Request) {
	// Define the relative path from the project root to the template file
	relativePath := "../../web/templates/worktime.html"

	// Get the absolute path to the template
	tmplPath, err := filepath.Abs(relativePath)
	if err != nil {
		http.Error(w, "Failed to construct template path", http.StatusInternalServerError)
		return
	}

	// Rest of your handler logic...
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to execute template: "+err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		// Process POST request...
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// processWorkTimeForm processes the POST request with form data
func processWorkTimeForm(w http.ResponseWriter, r *http.Request) {
	// Make sure to call ParseForm() before accessing form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	startzeit := r.FormValue("startzeit")
	dauer, err := strconv.Atoi(r.FormValue("dauer")) // convert to int
	if err != nil {
		http.Error(w, "Invalid duration value", http.StatusBadRequest)
		return
	}

	pause, err := strconv.Atoi(r.FormValue("pause")) // convert to int
	if err != nil {
		http.Error(w, "Invalid break duration value", http.StatusBadRequest)
		return
	}

	// Parse start time
	startTime, err := time.Parse("15:04", startzeit)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}

	// Calculate end time
	endTime := startTime.Add(time.Minute * time.Duration(dauer+pause))

	// Send the result back to the client
	fmt.Fprintf(w, "End time of work: %s", endTime.Format("15:04"))
}

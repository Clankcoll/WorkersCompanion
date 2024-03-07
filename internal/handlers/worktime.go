package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

// Define a struct to hold the template data
type PageData struct {
	Result string
}

func WorkTimeCalculator(w http.ResponseWriter, r *http.Request) {
	tmplPath, err := filepath.Abs("../../web/templates/worktime.html")
	if err != nil {
		http.Error(w, "Failed to construct template path", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize an empty result
	data := PageData{Result: ""}

	if r.Method == "POST" {
		// Parse form values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
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

		// Update the result in the PageData
		data.Result = "End time of work: " + endTime.Format("15:04")
	}

	// Render the template with the PageData (which includes the result if POST was processed)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

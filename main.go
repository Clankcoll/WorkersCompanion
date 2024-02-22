package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	startzeit := r.FormValue("startzeit")
	dauerStr := r.FormValue("dauer")
	pauseStr := r.FormValue("pause")

	dauer, err := strconv.Atoi(dauerStr)
	if err != nil {
		http.Error(w, "Dauer muss eine Zahl sein", http.StatusBadRequest)
		return
	}

	pause, err := strconv.Atoi(pauseStr)
	if err != nil {
		http.Error(w, "Pause muss eine Zahl sein", http.StatusBadRequest)
		return
	}

	start, err := time.Parse("15:04", startzeit)
	if err != nil {
		http.Error(w, "Ungültiges Zeitformat", http.StatusBadRequest)
		return
	}

	// Berechnung der Endzeit ohne Pause
	endeOhnePause := start.Add(time.Minute * time.Duration(dauer))

	// Berechnung der tatsächlichen Endzeit mit Pause
	ende := endeOhnePause.Add(time.Minute * time.Duration(pause))

	ergebnis := "Das Ende ist um " + ende.Format("15:04") + " Uhr."

	w.Write([]byte(ergebnis))
}

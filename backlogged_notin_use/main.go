package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	lang := getPreferredLanguage(r)
	var tmpl *template.Template
	var err error

	// Überprüfe die Sprache des Browsers
	if lang == "en" {
		tmpl, err = template.ParseFiles("index_en.html")
	} else {
		tmpl, err = template.ParseFiles("index.html")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func getPreferredLanguage(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")
	langs := strings.Split(lang, ",")
	for _, l := range langs {
		l = strings.Split(l, ";")[0]
		l = strings.TrimSpace(l)
		if l == "en" || l == "de" { // hier mehr sprachen hinzufügen
			return strings.ToLower(l)
		}
	}
	return "" // kehre zur default language zurück when keine sprache gefunden
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

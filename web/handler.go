package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Text   string
	Banner string
	Result string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	templ.Execute(w, PageData{})
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "field cannot be empty", http.StatusBadRequest)
		return
	}
	m, err := LoadBnner(banner)
	if err != nil {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}
	err = validate(text)
	if err != nil {
		http.Error(w, "bad reqquest", http.StatusBadRequest)
		return
	}
	result, err := Generate(text, m)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	templ.Execute(w, PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	})
}
func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/switch-handler" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}
}

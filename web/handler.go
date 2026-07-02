package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
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

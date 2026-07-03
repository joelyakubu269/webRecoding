package main

import "net/http"

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art-switch", SwitchHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.ListenAndServe(":8080", nil)
}

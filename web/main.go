package main

import "net/http"

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/switch-handler", SwitchHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.ListenAndServe(":8080", nil)
}

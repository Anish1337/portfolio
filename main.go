package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
}

// Global Scope
// home function handles request to homepage
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")

}

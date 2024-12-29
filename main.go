package main

import (
    "fmt"
    "html/template"
    "log"
	"net/http"
)

func main() {
    // route handler
	http.HandleFunc("/", home)
    http.HandleFunc("/projects", projects)
    http.HandleFunc("/skills", skills)
    http.HandleFunc("/about", about)
    http.HandleFunc("/contact", contact)

    // start server
    htpp.ListenAndServer(":8008", nil)
}

// Global Scope

// home function handles request to homepage
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")

    
}  

func rendedrTempalte(w http.ResponseWrioter, tmpl string){
    t, err := template.ParseFiles("tempaltes/" + tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
    }
    t.Execute(w, nil)
}

package main // Declare the main package

import (
	"fmt"           // Import the fmt package for formatted I/O
	"html/template" // Import the html/template package for rendering HTML templates
	"log"           // Import the log package for logging errors
	"net/http"      // Import the net/http package for HTTP server and client
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/skills", skills)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("🚀 Server is running on http://localhost:8080")

    // Start server (port 8080)
	err := http.ListenAndServe(":8080", nil)
   
    // Check + log errors
	if err != nil { 
		log.Fatal(err) 
	}
}

// page functions

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html") 
}

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html") 
}

func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html") 
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html") 
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html") 
}

// renderTemplate function
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl) 

    // error check
	if err != nil {                                    
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

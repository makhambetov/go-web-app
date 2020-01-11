package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>Home page</>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "text/html")
	fmt.Fprint(w, "<h1>Contact page</>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry. Page not found</>")
}


func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	fmt.Println("listening 3000")
	http.ListenAndServe(":3000", r)
}

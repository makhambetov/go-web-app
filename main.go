package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello!</h1>")

}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening 3000")
	http.ListenAndServe(":3000", nil)
}

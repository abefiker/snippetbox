package main

import (
	"log"
	"net/http"
	"strconv"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Create a new snippet"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server at on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

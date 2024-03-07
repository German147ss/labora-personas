package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	mux.HandleFunc("POST /task/{id}/details", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	http.ListenAndServe(":8080", mux)
}

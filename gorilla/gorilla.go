package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

func main() {
	r := mux.NewRouter()
	apmgorilla.Instrument(r)

	r.HandleFunc("/books/{title}/page/{page}", serveBookAndPage)

	log.Printf("Listening on localhost")

	_ = http.ListenAndServe(":80", r)
}

func serveBookAndPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	_, _ = fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

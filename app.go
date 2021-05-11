package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andreanidouglas/url-shortener/handlers"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	l := log.New(os.Stdout, "links-api ", log.LstdFlags)

	r := mux.NewRouter()
	lh := handlers.NewLink(l)

	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/links", lh.GetAllLinks).Methods("GET")
	r.HandleFunc("/links", lh.PostLink).Methods("POST")

	fmt.Println("Serving on port: 3000")
	http.ListenAndServe(":3000", r)

}

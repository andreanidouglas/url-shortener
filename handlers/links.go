package handlers

import (
	"log"
	"net/http"

	"github.com/andreanidouglas/url-shortener/data"
)

type Links struct {
	l *log.Logger
}

func NewLink(l *log.Logger) *Links {
	return &Links{l}
}

func (l *Links) PostLink(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	link := &data.Link{}

	err := link.FromJson(r.Body)
	if err != nil {
		http.Error(w, "could not unmarshall data", http.StatusBadRequest)
		l.l.Fatalf("could not unmarshall data properly: %v", err)
		return
	}

	data.AddLink(link)

}

func (l *Links) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	links := data.GetLinks()

	err := links.ToJson(w)
	if err != nil {
		http.Error(w, "could not get links", http.StatusInternalServerError)
		l.l.Fatalf("could not marshall data properly: %v", err)
		return
	}

}

func setupHeaders(w http.ResponseWriter) {

	w.Header().Add("content-type", "application/json")

}

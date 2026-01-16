package handlers

import "net/http"

type Shortener struct {
	// Deps/?
}

func NewShortener() *Shortener {
	return &Shortener{}
}

func (h *Shortener) ShortenHandler(w http.ResponseWriter, r *http.Request) {

}

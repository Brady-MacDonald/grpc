package handlers

import "net/http"

type Redirects struct {
}

func NewRedirects() *Redirects {
	return &Redirects{}
}

func (h *Redirects) RedirectHandler(w http.ResponseWriter, r *http.Request) {}

package handlers

import "net/http"

type handler struct {
	Shorten  http.Handler
	Redirect http.Handler
}

func NewHandler() *handler {
	return &handler{}
}

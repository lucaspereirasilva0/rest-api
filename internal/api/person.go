package api

import "net/http"

type Handler interface {
	GetPerson(w http.ResponseWriter, r *http.Request)
}
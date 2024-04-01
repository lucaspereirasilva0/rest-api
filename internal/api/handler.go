package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucaspereirasilva0/rest-api/internal/business"
	"github.com/lucaspereirasilva0/rest-api/tools"
)

type handler struct{
	personService business.Service
}

func NewHandler(personService business.Service) Handler {
	return &handler{}
}

func (h *handler) GetPerson (w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all items..")
	personList, err := h.personService.GetPerson()
	if err != nil {
		tools.ApiEncode(w, err)
	}
	tools.ApiEncode(w, personList)
}

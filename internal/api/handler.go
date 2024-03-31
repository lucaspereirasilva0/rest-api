package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucaspereirasilva0/rest-api/internal/business"
	"log"
	"net/http"
)

func StartHandler() {
	router := chi.NewRouter()
	router.Get("/contato", business.GetPerson)
	router.Get("/contato/{id}", business.GetPersonId)
	router.Post("/contato", business.CreatePerson)
	router.Post("/contato/{filename}", business.CreatePersonFromFile)
	router.Put("/contato", business.UpdatePerson)
	router.Delete("/contato/{id}", business.DeletePerson)
	log.Println(http.ListenAndServe(":8085", router))
}

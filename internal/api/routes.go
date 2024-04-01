package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartHandler(handler Handler) {
	router := chi.NewRouter()
	router.Get("/contato", handler.GetPerson)
/* 	router.Get("/contato/{id}", business.GetPersonId)
	router.Post("/contato", business.CreatePerson)
	router.Post("/contato/{filename}", business.CreatePersonFromFile)
	router.Put("/contato", business.UpdatePerson)
	router.Delete("/contato/{id}", business.DeletePerson) */
	log.Println(http.ListenAndServe(":8085", router))
}
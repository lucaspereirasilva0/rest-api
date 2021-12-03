package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucaspereirasilva0/rest-api/internal/business"
	"log"
	"net/http"
)

func StartHandler() {
	router := chi.NewRouter()
	router.Get("/findperson", business.GetPerson)
	router.Get("/findperson/{id}", business.GetPersonId)
	router.Post("/createperson", business.CreatePerson)
	router.Post("/createpersonfile/{filename}", business.CreatePersonFromFile)
	router.Put("/updateperson", business.UpdatePerson)
	router.Delete("/deleteperson/{id}", business.DeletePerson)
	log.Println(http.ListenAndServe(":8080", router))
}

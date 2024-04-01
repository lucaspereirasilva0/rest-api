package main

import (
	"fmt"

	"github.com/lucaspereirasilva0/rest-api/internal/api"
	"github.com/lucaspereirasilva0/rest-api/internal/business"
	"github.com/lucaspereirasilva0/rest-api/internal/repository"
	"github.com/lucaspereirasilva0/rest-api/internal/scripts"
)

func main() {
	//scripts.LoadPerson()
	repositoryService, err := repository.NewRepository()
	if err != nil {
		panic("error create repository service: " + err.Error())
	}
	personService := business.NewService(repositoryService)
	handler := api.NewHandler(personService)
	api.StartHandler(handler)
}

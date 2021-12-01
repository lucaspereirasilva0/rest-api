package main

import (
	"github.com/lucaspereirasilva0/rest-api/internal/api"
	"github.com/lucaspereirasilva0/rest-api/internal/scripts"
)

func main() {
	scripts.LoadPerson()
	api.StartHandler()
}

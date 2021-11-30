package main

import (
	"github.com/go-chi/chi/v5"
	//"fmt"
	"log"
	"net/http"
	"rest-api/internal/business"
)

func main() {
	router := chi.NewRouter()
	router.Get("/contato", business.GetPeople)
	router.Get("/contato/{id}", business.GetPerson)
	router.Post("/contato", business.CreatePerson)
	router.Post("/contato/{filename}", business.CreatePersonFromFile)
	router.Put("/contato", business.UpdatePerson)
	router.Delete("/contato/{id}", business.DeletePerson)
	router.Delete("/deletetable", business.DeleteTable)
	router.Post("/createtable", business.CreateTable)
	log.Fatal(http.ListenAndServe(":8080", router))

	//for i := 0; i < 3; i++  {
	//	persons := createSliceManually(strconv.Itoa(i))
	//	for _, person := range persons {
	//		fmt.Println("Put a item..")
	//	    fmt.Println(person)
	//		putItem(OpenDynamoDBLocal(), person)
	//	}
	//}

	// fmt.Println("Creating tables..")
	// createTable(OpenDynamoDBLocal())

	//fmt.Println("Deleting tables..")
	//deleteTable(OpenDynamoDBLocal())

	//fmt.Println("Listing tables..")
	//listTables(OpenDynamoDBLocal())

	//fmt.Println("Getting items..")
	//getAllItems(OpenDynamoDBLocal())

	//fmt.Println("Getting a item..")
	//getItem(OpenDynamoDBLocal(),"1")

}

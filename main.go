package main

import (
	//"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/contato", GetPeople)
	router.Get("/contato/{id}", GetPerson)
	router.Post("/contato", CreatePerson)
	router.Post("/contato/{filename}", CreatePersonFromFile)
	router.Put("/contato", UpdatePerson)
	router.Delete("/contato/{id}", DeletePerson)
	router.Delete("/deletetable", DeleteTable)
	router.Post("/createtable", CreateTable)
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

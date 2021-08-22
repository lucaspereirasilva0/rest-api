package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
)

type Person struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	//Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func OpenDynamoDBLocal() *dynamodb.Client{
	svc , err := loadDatabase()
	if err != nil {
		e := apiErrors("fail to load database", err)
		log.Println(e)
		apiEncode(nil, e)
		return nil
	}

	return svc
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all items..")
	person, err := getAllItems(OpenDynamoDBLocal())
	if err != nil {
		e := apiErrors("fail to get all items", err)
		log.Println(e)
		apiEncode(w, e)
		return
	}

	apiEncode(w, person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a item...")

	person, err := getItem(OpenDynamoDBLocal(), chi.URLParam(r, "id"))
	if err != nil {
		e := apiErrors("fail to get an item", err)
		log.Println(e)
		apiEncode(w, e)
		return
	} else {
		if person == (Person{}) { 
			log.Println("person not found")
			apiEncode(w, "person not found")
			return
		}
	}

	apiEncode(w, person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person

	log.Println("Putting an item")

	apiDecode(r, &person)

	if person == (Person{}) {
		e := apiErrors("person nil, fail in decode", fmt.Errorf("see the log"))
		log.Println(e)
		apiEncode(w, e)
		return
	}

	svc, errOpenDB := loadDatabase()
	if errOpenDB != nil {
		e := apiErrors("fail to open database", errOpenDB)
		log.Println(e)
		apiEncode(w, e)
		return
	}

	persons, errGetAllItems := getAllItems(svc)

	if errGetAllItems != nil {
		e := apiErrors("fail to get all items", errGetAllItems)
		log.Println(e)
		apiEncode(w, e)
		return
	}

	id := len(persons) + 1

	person.ID = id

	err := putItem(OpenDynamoDBLocal(), person)

	if err != nil {
		e := apiErrors("fail to put a item", err)
		log.Println(e)
		apiEncode(w, e)
		return
	} else {
		apiEncode(w, person)
		apiEncode(w, "put an item success")
		log.Println("put an item success")
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting an item")

	err := deleteItem(OpenDynamoDBLocal(), chi.URLParam(r, "id"))
	if err != nil {
		e := apiErrors("fail to delete a item", err)
		log.Println(e)
		apiEncode(w, e)
		return
	}else {
		log.Println("delete item success")
		apiEncode(w, "delete item success")
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person

	log.Println("Update a item")

	errDecode := json.NewDecoder(r.Body).Decode(&person)
	if errDecode != nil {
		e := apiErrors("fail to decode json to struct", errDecode)
		log.Println(e)
		apiEncode(w, e)
		return
	}


	err := putItem(OpenDynamoDBLocal(), person)
	if err != nil {
		e := apiErrors("fail to put an item", err)
		log.Println(e)
		apiEncode(w, e)
		return
	}else{
		apiEncode(w, person)
		log.Println("update item success")
	}
}

func DeleteTable(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a table...")
	deleteTable(OpenDynamoDBLocal())
	log.Println("delete table success")
	apiEncode(w, "delete table success")
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a table...")
	createTable(OpenDynamoDBLocal())
	log.Println("create table success")
	apiEncode(w, "create table success")
}

func (p Person) saveToFile() error {
	file, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		e := apiErrors("fail in format file to json", err)
		log.Println(e)
		return err
	}
	_ = ioutil.WriteFile("persons.json", file, 0666)

	return nil
}

func CreatePersonFromFile(w http.ResponseWriter, r *http.Request) {
	var persons []PersonDynamo

	log.Println("Putting an item from file")

	jsonFile, errOpenFile := os.Open("persons.json")
	if errOpenFile != nil {
		e := apiErrors("fail to open a json file", errOpenFile)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
		return
	}

	defer jsonFile.Close()

	b, errReadAll := ioutil.ReadAll(jsonFile)
	if errReadAll != nil {
		e := apiErrors("fail to read a json file", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
		return
	}

	errUnmarshal := json.Unmarshal(b, &persons)
	if errUnmarshal != nil {
		e := apiErrors("fail in unmarshal a json file to person", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
		return
	}

	for _, p := range persons {
		person := Person{
			ID:        p.Id,
			Firstname: p.FirstName,
			Lastname:  p.LastName,
		}

		err := putItem(OpenDynamoDBLocal(), person)
		if err != nil {
			e := apiErrors("fail to put a item", err)
			log.Println(e)
			//_ = json.NewEncoder(w).Encode(_createPerson)
			return
		} else {
			_ = json.NewDecoder(r.Body).Decode(&person)
			_ = json.NewEncoder(w).Encode(person)
			log.Println("put an item success")
			_ = json.NewEncoder(w).Encode("put an item success")
		}
	}

}

//func readFile() error {
//	var person Person
//	file, _ := ioutil.ReadAll("persons.json")
//	json.Unmarshal(file, &person)
//}

//func createSliceManually(id string) []Person{
//
//	if id == "1" {
//		people = append(people, Person{
//			ID:        id,
//			Firstname: "Lucas",
//			Lastname:  "Pereira",
//			//Address:   &Address{
//			//	City:  "São Paulo",
//			//	State: "São Paulo",
//			//},
//		})
//	}else {
//		if id == "2" {
//			people = append(people, Person{
//				ID:        id,
//				Firstname: "Joao",
//				Lastname:  "Souza",
//				//Address:   &Address{
//				//	City:  "Extrema",
//				//	State: "Minas Gerais",
//				//},
//			})
//		}else {
//			people = append(people, Person{
//				ID:        id,
//				Firstname: "Elias",
//				Lastname:  "Reis",
//				//Address:   &Address{
//				//	City:  "Rio de Janeiro",
//				//	State: "Rio de Janeiro",
//				//},
//			})
//		}
//	}
//
//	return people
//}

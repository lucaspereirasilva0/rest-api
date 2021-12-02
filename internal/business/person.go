package business

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lucaspereirasilva0/rest-api/internal/errors"
	"github.com/lucaspereirasilva0/rest-api/internal/model"
	"github.com/lucaspereirasilva0/rest-api/internal/repositories"
	"github.com/lucaspereirasilva0/rest-api/internal/repositories/person"
	"github.com/lucaspereirasilva0/rest-api/tools"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func OpenConnection() *dynamodb.Client {
	svc, err := repositories.NewRepository()
	if err != nil {
		e := errors.New("fail to load database", err)
		log.Println(e)
		tools.ApiEncode(nil, e)
		return nil
	}
	return svc
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all items..")
	p, err := person.GetAllItems(OpenConnection(), "person")
	if err != nil {
		e := errors.New("fail to get all items", err)
		log.Println(e)
		tools.ApiEncode(w, e)
		return
	}
	tools.ApiEncode(w, p)
}

func GetPersonId(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a item...")

	p, err := person.GetItem(OpenConnection(), chi.URLParam(r, "id"))
	if err != nil {
		e := errors.New("fail to get an item", err)
		log.Println(e)
		tools.ApiEncode(w, e)
	} else {
		if p == (model.Person{}) {
			log.Println("person not found")
			tools.ApiEncode(w, "person not found")
		}
	}

	tools.ApiEncode(w, p)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person

	log.Println("Putting an item")
	tools.ApiDecode(r, &p)

	if p == (model.Person{}) {
		e := errors.New("person nil, fail in decode", fmt.Errorf("see the log"))
		log.Println(e)
		tools.ApiEncode(w, e)
	}

	svc := OpenConnection()

	p.Id = uuid.NewString()

	err := person.PutItem(svc, p, "person")

	if err != nil {
		e := errors.New("fail to put a item", err)
		log.Println(e)
		tools.ApiEncode(w, e)
	} else {
		tools.ApiEncode(w, p)
		log.Println("put an item success")
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting an item")

	err := person.DeleteItem(OpenConnection(), chi.URLParam(r, "id"))
	if err != nil {
		e := errors.New("fail to delete a item", err)
		log.Println(e)
		tools.ApiEncode(w, e)
	} else {
		log.Println("delete item success")
		tools.ApiEncode(w, "delete item success")
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person

	log.Println("Update a item")

	errDecode := json.NewDecoder(r.Body).Decode(&p)
	if errDecode != nil {
		e := errors.New("fail to decode json to struct", errDecode)
		log.Println(e)
		tools.ApiEncode(w, e)
	}

	err := person.PutItem(OpenConnection(), p, "person")
	if err != nil {
		e := errors.New("fail to put an item", err)
		log.Println(e)
		tools.ApiEncode(w, e)
	} else {
		tools.ApiEncode(w, p)
		log.Println("update item success")
	}
}

//func (p Person) saveToFile() error {
//	file, err := json.MarshalIndent(p, "", " ")
//	if err != nil {
//		e := errors.New("fail in format file to json", err)
//		log.Println(e)
//		return err
//	}
//	_ = ioutil.WriteFile("persons.json", file, 0666)
//
//	return nil
//}

func CreatePersonFromFile(w http.ResponseWriter, r *http.Request) {
	var persons []person.Person

	log.Println("Putting an item from file")

	jsonFile, errOpenFile := os.Open("persons.json")
	if errOpenFile != nil {
		e := errors.New("fail to open a json file", errOpenFile)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
	}

	defer jsonFile.Close()

	b, errReadAll := ioutil.ReadAll(jsonFile)
	if errReadAll != nil {
		e := errors.New("fail to read a json file", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
	}

	errUnmarshal := json.Unmarshal(b, &persons)
	if errUnmarshal != nil {
		e := errors.New("fail in unmarshal a json file to person", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
	}

	for _, p := range persons {
		p := model.Person{
			Id:        p.Id,
			Firstname: p.FirstName,
			Lastname:  p.LastName,
		}

		err := person.PutItem(OpenConnection(), p, "person")
		if err != nil {
			e := errors.New("fail to put a item", err)
			log.Println(e)
			//_ = json.NewEncoder(w).Encode(_createPerson)
		} else {
			_ = json.NewDecoder(r.Body).Decode(&p)
			_ = json.NewEncoder(w).Encode(p)
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

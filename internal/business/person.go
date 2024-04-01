package business

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lucaspereirasilva0/rest-api/internal/errors"
	"github.com/lucaspereirasilva0/rest-api/internal/model"
	"github.com/lucaspereirasilva0/rest-api/internal/repository"
	"github.com/lucaspereirasilva0/rest-api/tools"
	"log"
	"net/http"
	"os"
)

type Service interface {
	GetPerson() ([]model.Person, error)
}

/* func OpenDynamoDBLocal() *dynamodb.Client {
	svc, err := person.LoadDatabase()
	if err != nil {
		e := errors.New("fail to load database", err)
		log.Println(e)
		tools.ApiEncode(nil, e)
		return nil
	}

	return svc
} */

/* func GetPersonId(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a item...")

	p, err := person.GetItem(OpenDynamoDBLocal(), chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		tools.ApiEncode(w, NewGetItemError(err))
		return
	} else {
		if p == (model.Person{}) {
			tools.ApiEncode(w, NewPersonNotFoundError())
			return
		}
	}

	tools.ApiEncode(w, p)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person

	log.Println("Putting an item")
	tools.ApiDecode(r, &p)

	if p == (model.Person{}) {
		tools.ApiEncode(w, NewFailDecodeError(nil))
		return
	}

	svc := OpenDynamoDBLocal()

	p.Id = uuid.NewString()

	err := person.PutItem(svc, p, "person")

	if err != nil {
		log.Println(err)
		tools.ApiEncode(w, NewPutItemError(err))
		return
	} else {
		tools.ApiEncode(w, p)
		log.Println("put an item success")
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting an item")

	err := person.DeleteItem(OpenDynamoDBLocal(), chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		tools.ApiEncode(w, NewDeleteItemError(err))
		return
	} else {
		log.Println("delete item success")
		tools.ApiEncode(w, "delete item success")
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person

	log.Println("Update a item")

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		tools.ApiEncode(w, NewFailDecodeError(err))
		return
	}

	err = person.PutItem(OpenDynamoDBLocal(), p, "person")
	if err != nil {
		log.Println(err)
		tools.ApiEncode(w, NewPutItemError(err))
		return
	} else {
		tools.ApiEncode(w, p)
		log.Println("update item success")
	}
} */

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

/* func CreatePersonFromFile(w http.ResponseWriter, r *http.Request) {
	var persons []person.Person

	log.Println("Putting an item from file")

	b, errReadAll := os.ReadFile("persons.json")
	if errReadAll != nil {
		e := errors.New("fail to read a json file", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
		return
	}

	errUnmarshal := json.Unmarshal(b, &persons)
	if errUnmarshal != nil {
		e := errors.New("fail in unmarshal a json file to person", errReadAll)
		log.Println(e)
		//_ = json.NewEncoder(w).Encode(_createPersonFromFile)
		return
	}

	for _, p := range persons {
		p := model.Person{
			Id:        p.Id,
			Firstname: p.FirstName,
			Lastname:  p.LastName,
		}

		err := person.PutItem(OpenDynamoDBLocal(), p, "person")
		if err != nil {
			e := errors.New("fail to put a item", err)
			log.Println(e)
			//_ = json.NewEncoder(w).Encode(_createPerson)
			return
		} else {
			_ = json.NewDecoder(r.Body).Decode(&p)
			_ = json.NewEncoder(w).Encode(p)
			log.Println("put an item success")
			_ = json.NewEncoder(w).Encode("put an item success")
		}
	}

} */

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

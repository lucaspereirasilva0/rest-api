package scripts

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lucaspereirasilva0/rest-api/internal/errors"
	"github.com/lucaspereirasilva0/rest-api/internal/repositories"
	"github.com/lucaspereirasilva0/rest-api/internal/repositories/person"
	"log"
)

func openDynamoDBLocal() *dynamodb.Client {
	svc, err := repositories.NewRepository()
	if err != nil {
		e := errors.New("fail to load database", err)
		log.Println(e)
		return nil
	}
	return svc
}

func LoadPerson() {
	p := "person"
	svc := openDynamoDBLocal()
	log.Println("Creating a table...")
	tables, errList := person.ListTables(svc)
	if errList != nil {
		e := errors.New("fail to list tables", errList)
		log.Println(e)
	}
	if len(tables) != 0 {
		for _, t := range tables {
			if t != p {
				createTable(svc, p)
			}
		}
	} else {
		createTable(svc, p)
	}
}

func createTable(svc *dynamodb.Client, tableName string) {
	err := person.CreateTable(svc, tableName)
	if err != nil {
		e := errors.New("fail to create table", err)
		log.Println(e)
	} else {
		log.Println("create table success")
	}
}

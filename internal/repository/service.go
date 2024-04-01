package repository

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lucaspereirasilva0/rest-api/internal/errors"
	"github.com/lucaspereirasilva0/rest-api/internal/model"
)

type service struct {
	dbClient *dynamodb.Client
}

func NewRepository() (Service, error) {
	client, err := NewDynamoClient()
	if err != nil {
		return nil, err
	}
	return &service{
		dbClient: client,
	}, nil
}

func (s *service) GetAllItems() ([]model.Person, error) {
	var person []model.Person

	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("person"),
	}

	result, err := s.dbClient.Scan(context.TODO(), scanInput)

	if err != nil {
		e := errors.New("fail to get all items", err)
		log.Println(e)
		return []model.Person{}, err
	}

	errUnmarshal := attributevalue.UnmarshalListOfMaps(result.Items, &person)
	if errUnmarshal != nil {
		e := errors.New("fail to unmarshal items", errUnmarshal)
		log.Println(e)
		return []model.Person{}, err
	}

	log.Println(person)

	return person, nil
}


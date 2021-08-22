package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PersonDynamo struct {
	Id        int    `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstname"`
	LastName  string `dynamodbav:"lastname"`
}

type Id struct {
	Id int `dynamodbav:"id"`
}

func listTables(svc *dynamodb.Client) {
	params := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(context.TODO(), params)

	if err != nil {
		log.Println("error in ListTables of listTables", err)
	}

	log.Println("found the tables: ", result.TableNames)
}

func createTable(svc *dynamodb.Client) {
	createTableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String("my-table"),
		BillingMode: types.BillingModePayPerRequest,
	}
	_, err := svc.CreateTable(context.TODO(), createTableInput)
	if err != nil {
		log.Println("error in CreateTable of createTable: ", err)
	}

	log.Println(createTableInput.TableName, "", createTableInput.AttributeDefinitions)
}

func loadDatabase() (*dynamodb.Client, error) {
	os.Setenv("AWS_ACCESS_KEY_ID", "dummy1")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy2")
	os.Setenv("AWS_SESSION_TOKEN", "dummy3")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL("http://localhost:8000")
	})

	return svc, nil
}

func deleteTable(svc *dynamodb.Client) {
	deleteTableInput := &dynamodb.DeleteTableInput{
		TableName: aws.String("my-table"),
	}
	_, err := svc.DeleteTable(context.TODO(), deleteTableInput)
	if err != nil {
		log.Println("error in DeleteTable of deleteTable: ", err)
	}

	log.Println(deleteTableInput.TableName)
}

func getAllItems(svc *dynamodb.Client) ([]Person, error) {
	var person []Person

	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("my-table"),
	}

	result, err := svc.Scan(context.TODO(), scanInput)

	if err != nil {
		e := apiErrors("fail to get all items", err)
		log.Println(e)
		//return person, e
	}

	errUnmarshal := attributevalue.UnmarshalListOfMaps(result.Items, &person)
	if errUnmarshal != nil {
		e := apiErrors("fail to unmarshal items", errUnmarshal)
		log.Println(e)
		//return person, e
	}

	for _, p := range person {
		log.Println(p)
	}

	return person, nil
}

func getAllItemsWithCondition(svc *dynamodb.Client) {
	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        aws.String("my-table"),
		FilterExpression: aws.String("attribute_not_exists(deletedAt) AND contains(firstName, :firstName)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":firstName": &types.AttributeValueMemberS{Value: "John"},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Items)
}

func getAllItemsWithConditionExpressions(svc *dynamodb.Client) {
	expr, err := expression.NewBuilder().WithFilter(
		expression.And(
			expression.AttributeNotExists(expression.Name("deletedAt")),
			expression.Contains(expression.Name("firstName"), "John"),
		),
	).Build()
	if err != nil {
		panic(err)
	}

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String("my-table"),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out.Items)
}

func getItem(svc *dynamodb.Client, id string) (Person, error) {

	var person Person

	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{
				Value: id,
			},
		},
		TableName: aws.String("my-table"),
	}

	result, err := svc.GetItem(context.TODO(), getItemInput)

	if err != nil {
		return person, err
	}

	errUnmarshalMap := attributevalue.UnmarshalMap(result.Item, &person)
	if errUnmarshalMap != nil {
		log.Println("error in UnmarshalMap &person of getItem: ", errUnmarshalMap)
	}

	log.Println(person)

	return person, nil
}

func putItem(svc *dynamodb.Client, param Person) error {
	persons := PersonDynamo{
		Id:        param.ID,
		FirstName: param.Firstname,
		LastName:  param.Lastname,
	}

	item, errMarshalMap := attributevalue.MarshalMap(persons)
	if errMarshalMap != nil {
		return errMarshalMap
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("my-table"),
		Item:      item,
	}
	_, err := svc.PutItem(context.TODO(), putItemInput)
	if err != nil {
		return err
	}

	var person Person
	errUnmarshalMap := attributevalue.UnmarshalMap(putItemInput.Item, &person)
	if errUnmarshalMap != nil {
		return errUnmarshalMap
	}

	log.Println(person)

	return nil
}

func deleteItem(svc *dynamodb.Client, id string) error {
	deleteItemInput := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
		},
		TableName: aws.String("my-table"),
	}

	_, err := svc.DeleteItem(context.TODO(), deleteItemInput)

	if err != nil {
		return err
	}

	return nil
}

//func listTablesOld(svc *dynamodb.Client)  {
//	params := &dynamodb.ListTablesInput{}
//
//	result, err := svc.ListTables(context.TODO(), params)
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(result.TableNames)
//
//	//p := dynamodb.NewListTablesPaginator(svc, nil, func(o *dynamodb.ListTablesPaginatorOptions) {
//	//	o.StopOnDuplicateToken = true
//	//})
//
//	//for p.HasMorePages() {
//	//	out, err := p.NextPage(context.TODO())
//	//	if err != nil {
//	//		panic(err)
//	//	}
//
//	//	for _, tn := range out.TableNames {
//	//		fmt.Println(tn)
//	//	}
//	//}
//}

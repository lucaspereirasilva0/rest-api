package repository

type Person struct {
	Id        string `dynamodbav:"id"`
	FirstName string `dynamodbav:"firstname"`
	LastName  string `dynamodbav:"lastname"`
}

type Address struct {
	Id    string `dynamodbav:"id"`
	City  string `dynamodbav:"city"`
	State string `dynamodbav:"state"`
}

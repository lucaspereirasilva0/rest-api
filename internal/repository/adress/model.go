package adress

type Address struct {
	Id    string `dynamodbav:"id"`
	City  string `dynamodbav:"city"`
	State string `dynamodbav:"state"`
}

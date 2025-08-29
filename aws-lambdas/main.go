package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is you name?"`
	Age  int    `json:"what is your age?"`
}

type MyResponse struct {
	Message string `json:"answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{
		Message: fmt.Sprintf("%s is %d year old", event.Name, event.Age),
	}, nil
}
func main() {
	lambda.Start(HandleLambdaEvent)
}

package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(processRequest)
}

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func processRequest(r Request) (Response, error) {
	log.Println("Request =>")
	log.Print(r)

	return Response{
		StatusCode: http.StatusOK,
		Body:       string("{ \"message\": \"Hello\"}"),
	}, nil
}

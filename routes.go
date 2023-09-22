package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var validate *validator.Validate = validator.New()
var resume Resume

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received req %#v", req)

	switch req.HTTPMethod {
	case "GET":
		return processGetResume(ctx, req)
	case "POST":
		return processUpdateResume(ctx, req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func processGetResume(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Received GET resume request")

	// Simulate fetching the resume data
	jsonResume, err := json.Marshal(resume)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonResume),
	}, nil
}

func processUpdateResume(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Received PUT resume request")

	// Parse the request body into a Resume struct
	var updatedResume Resume
	err := json.Unmarshal([]byte(req.Body), &updatedResume)
	if err != nil {
		return clientError(http.StatusBadRequest)
	}

	// Validate the updated resume data
	err = validate.Struct(&updatedResume)
	if err != nil {
		return clientError(http.StatusBadRequest)
	}

	// Update the global resume data with the new values
	resume = updatedResume

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(status),
		StatusCode: status,
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
	}, nil
}

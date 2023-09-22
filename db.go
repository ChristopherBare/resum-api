package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	_ "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

const TableName = "Resume"

var db dynamodb.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.Resume())
	if err != nil {
		log.Fatal(err)
	}

	db = *dynamodb.NewFromConfig(sdkConfig)
}

type Resume struct {
	Id        string      `json:"id" dynamodbav:"id"`
	Education []Education `json:"education" dynamodbav:"education"`
	Jobs      []Job       `json:"jobs" dynamodbav:"jobs"`
	Projects  []Project   `json:"projects" dynamodbav:"projects"`
	Skills    []Skill     `json:"skills" dynamodbav:"skills"`
}

type Education struct {
	Id          string `json:"id" dynamodbav:"id"`
	School      string `json:"school" dynamodbav:"school"`
	Degree      string `json:"degree" dynamodbav:"degree"`
	AreaOfStudy string `json:"area-of-study" dynamodbav:"area-of-study"`
}

type Job struct {
	Id             string `json:"id" dynamodbav:"id"`
	Company        string `json:"company" dynamodbav:"company"`
	Title          string `json:"title" dynamodbav:"title"`
	JobDescription string `json:"job-description" dynamodbav:"job-description"`
	StartDate      string `json:"start-date" dynamodbav:"start-date"`
	EndDate        string `json:"end-date" dynamodbav:"end-date"`
}

type Project struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
}

type Skill struct {
	Id   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func saveResume(ctx context.Context, resume Resume) error {
	// Use the DynamoDB client to save the resume item
	av, err := dynamodbattribute.MarshalMap(resume)
	if err != nil {
		log.Printf("Error marshaling resume: %v", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &TableName,
	}

	_, err = db.PutItem(ctx, input)
	if err != nil {
		log.Printf("Error saving resume: %v", err)
		return err
	}

	return nil
}

func getResume(ctx context.Context, resumeID string) (Resume, error) {
	// Use the DynamoDB client to retrieve the resume item
	input := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: &resumeID,
			},
		},
		TableName: &TableName,
	}

	result, err := db.GetItem(ctx, input)
	if err != nil {
		log.Printf("Error fetching resume: %v", err)
		return Resume{}, err
	}

	if len(result.Item) == 0 {
		return Resume{}, fmt.Errorf("Resume not found")
	}

	var resume Resume
	err = dynamodbattribute.UnmarshalMap(result.Item, &resume)
	if err != nil {
		log.Printf("Error unmarshaling resume: %v", err)
		return Resume{}, err
	}

	return resume, nil
}

func main() {
	// Lambda function entry point
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// Your code here to handle API requests and interact with DynamoDB
	})
}

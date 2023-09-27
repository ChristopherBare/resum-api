package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const TableName = "Resume"

var db *dynamodb.Client

func init() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = dynamodb.NewFromConfig(cfg)
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
	// Marshal the Resume object into a DynamoDB AttributeValue map
	av, err := attributevalue.MarshalMap(resume)
	if err != nil {
		log.Printf("Error marshaling resume: %v", err)
		return err
	}

	// Create a PutItemInput with the marshaled Resume and table name
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	// Put the item into the DynamoDB table
	_, err = db.PutItem(ctx, input)
	if err != nil {
		log.Printf("Error saving resume: %v", err)
		return err
	}

	return nil
}

func getResume(ctx context.Context, resumeID string) (Resume, error) {
	// Create a GetItemInput with the key and table name
	input := &dynamodb.GetItemInput{
		Key: map[string]dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(resumeID),
			},
		},
		TableName: aws.String(TableName),
	}

	// Get the item from the DynamoDB table
	result, err := db.GetItem(ctx, input)
	if err != nil {
		log.Printf("Error fetching resume: %v", err)
		return Resume{}, err
	}

	// Check if the item was found
	if len(result.Item) == 0 {
		return Resume{}, fmt.Errorf("Resume not found")
	}

	// Unmarshal the DynamoDB item into a Resume object
	var resume Resume
	err = attributevalue.UnmarshalMap(result.Item, &resume)
	if err != nil {
		log.Printf("Error unmarshaling resume: %v", err)
		return Resume{}, err
	}

	return resume, nil
}

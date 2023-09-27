package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	dynamoDBClient *dynamodb.Client
	tableName      = "ResumeTable"
)

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

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("error loading AWS config: %v", err))
	}

	dynamoDBClient = dynamodb.NewFromConfig(cfg)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract the resume ID from the query string
	resumeID := request.QueryStringParameters["id"]
	if resumeID == "" {
		// If the 'id' parameter is not provided, perform a scan operation to retrieve all items from DynamoDB
		scanInput := &dynamodb.ScanInput{
			TableName: aws.String(tableName),
		}

		result, err := dynamoDBClient.Scan(ctx, scanInput)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprintf("Error: %v", err),
				StatusCode: 500,
			}, err
		}

		// Serialize the list of items to JSON
		var resumes []Resume
		for _, item := range result.Items {
			var resume Resume
			err := attributevalue.UnmarshalMap(item, &resume)
			if err != nil {
				return events.APIGatewayProxyResponse{
					Body:       fmt.Sprintf("Error: %v", err),
					StatusCode: 500,
				}, err
			}
			resumes = append(resumes, resume)
		}

		resumesJSON, err := json.Marshal(resumes)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       fmt.Sprintf("Error: %v", err),
				StatusCode: 500,
			}, err
		}

		// Customize the response body and status code
		response := events.APIGatewayProxyResponse{
			Body:       string(resumesJSON),
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return response, nil
	}

	// If the 'id' parameter is provided, retrieve the specific resume from DynamoDB
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: resumeID},
		},
	}

	result, err := dynamoDBClient.GetItem(ctx, getItemInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %v", err),
			StatusCode: 500,
		}, err
	}

	// Check if the resume was found in DynamoDB
	if result.Item == nil {
		return events.APIGatewayProxyResponse{
			Body:       "Resume not found",
			StatusCode: 404,
		}, nil
	}

	// Deserialize the resume data into the Resume struct
	var resume Resume
	err = attributevalue.UnmarshalMap(result.Item, &resume)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %v", err),
			StatusCode: 500,
		}, err
	}

	// Serialize the Resume struct to JSON
	resumeJSON, err := json.Marshal(resume)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %v", err),
			StatusCode: 500,
		}, err
	}

	// Customize the response body and status code
	response := events.APIGatewayProxyResponse{
		Body:       string(resumeJSON),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return response, nil
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

var (
	dynamoDBClient *dynamodb.Client
	tableName      = "ResumeTable"
)

type Resume struct {
	Id        string      `json:"id" dynamodbav:"id"`
	Name      string      `json:"name" dynamodbav:"name"`
	Education []Education `json:"education" dynamodbav:"education"`
	Jobs      []Job       `json:"jobs" dynamodbav:"jobs"`
	Projects  []Project   `json:"projects" dynamodbav:"projects"`
	Skills    []Skill     `json:"skills" dynamodbav:"skills"`
}

type Education struct {
	School      string `json:"school" dynamodbav:"school"`
	Degree      string `json:"degree" dynamodbav:"degree"`
	AreaOfStudy string `json:"area-of-study" dynamodbav:"area-of-study"`
}

type Job struct {
	Company        string `json:"company" dynamodbav:"company"`
	Title          string `json:"title" dynamodbav:"title"`
	JobDescription string `json:"job-description" dynamodbav:"job-description"`
	StartDate      string `json:"start-date" dynamodbav:"start-date"`
	EndDate        string `json:"end-date" dynamodbav:"end-date"`
}

type Project struct {
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
}

type Skill struct {
	Name string `json:"name" dynamodbav:"name"`
}

// UpdateRequest is used for PATCH requests to update a resume item.
type UpdateRequest struct {
	Id     string       `json:"id"`
	Update ResumeUpdate `json:"update"`
}

// ResumeUpdate represents the fields to be updated in a resume.
type ResumeUpdate struct {
	Id        *string     `json:"id,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Education []Education `json:"education,omitempty"`
	Jobs      []Job       `json:"jobs,omitempty"`
	Projects  []Project   `json:"projects,omitempty"`
	Skills    []Skill     `json:"skills,omitempty"`
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("error loading AWS config: %v", err))
	}

	dynamoDBClient = dynamodb.NewFromConfig(cfg)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handleGetRequest(ctx, request)
	case "POST":
		return handlePostRequest(ctx, request)
	case "PATCH":
		return handlePatch(ctx, request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "Unsupported HTTP method",
			StatusCode: 405,
		}, nil
	}
}

func handleGetRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "'*'",
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
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "'*'",
		},
	}

	return response, nil
}

func handlePostRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newResume Resume
	if err := json.Unmarshal([]byte(request.Body), &newResume); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error parsing request body: %v", err),
			StatusCode: 400,
		}, err
	}
	// Ensure the ID is unique
	newResume.Id = GenerateUniqueID()

	// Store the new resume in DynamoDB
	av, err := attributevalue.MarshalMap(newResume)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error marshaling resume data: %v", err),
			StatusCode: 500,
		}, err
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}
	fmt.Println(putItemInput)
	_, err = dynamoDBClient.PutItem(ctx, putItemInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error storing resume in DynamoDB: %v", err),
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Resume posted successfully",
		StatusCode: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func GenerateUniqueID() string {
	return uuid.NewString()
}

func handlePatch(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var updateRequest UpdateRequest
	if err := json.Unmarshal([]byte(request.Body), &updateRequest); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error parsing request body: %v", err),
			StatusCode: 400,
		}, err
	}

	// Check if the "Id" is provided in the request
	if updateRequest.Id == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Id is required for the update operation",
			StatusCode: 400,
		}, nil
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(
		expression.Set(expression.Name("Name"), expression.Value(updateRequest.Update.Name)),
		// Handle other fields (Education, Jobs, Projects, Skills) similarly
	).Build()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error building update expression: %v", err),
			StatusCode: 500,
		}, err
	}

	// Define the key of the item you want to update
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: updateRequest.Id},
	}

	updateItemInput := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       key,
		UpdateExpression:          updateExpression.Update(),
		ExpressionAttributeValues: updateExpression.Values(),
		ReturnValues:              types.ReturnValueAllNew,
	}

	_, err = dynamoDBClient.UpdateItem(ctx, updateItemInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error updating item: %v", err),
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Item updated successfully",
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

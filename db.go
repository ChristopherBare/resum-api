package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	_ "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

const TableName = "Resume"

var db dynamodb.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.RESUME())
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

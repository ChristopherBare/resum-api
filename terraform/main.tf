provider "aws" {
  region = "us-east-1"  # Update this to your desired AWS region
}

resource "aws_s3_bucket" "example" {
  bucket = "my-unique-bucket-name"  # Replace with your preferred bucket name
  acl    = "private"  # Adjust ACL as needed
}

module "my_module" {
  source = "./terraform"

  # Other module configuration options, if any
}

resource "aws_lambda_function" "resum-api" {
  function_name = "resum-api"
  handler      = "main"  # Update with your Lambda function's handler
  runtime      = "go1.8"  # Update with your Lambda function's runtime
  role         = aws_iam_role.lambda_role.arn

  environment {
    variables = {
      S3_BUCKET_NAME = aws_s3_bucket.example.bucket
    }
  }

  # Other Lambda function configuration options

  # Zip your Lambda function code
  filename = "path/to/your/lambda/function/code.zip"
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda-execution-role"

  # Attach policies as needed for your Lambda function
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

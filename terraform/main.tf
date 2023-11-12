locals {
  lambda_package_checksum = filebase64(var.lambda_package_path)
}

resource "random_id" "unique_id" {
  byte_length = 8 # You can adjust the length as needed
}

resource "aws_lambda_function" "resum_api_lambda" {
  function_name    = "resum-api-lambda-${var.github_repo}-${var.branch_name}"
  role             = aws_iam_role.resum_api_lambda_role.arn
  handler          = "main"
  runtime          = "go1.x"
  source_code_hash = local.lambda_package_checksum
  filename         = var.lambda_package_path
}

variable "lambda_package_path" {
  description = "Path to the Lambda deployment package"
  type        = string
  default     = "../lambda.zip"
}

resource "aws_iam_role" "resum_api_lambda_role" {
  name = "resum-api-lambda-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy" "resum_api_policy" {
  name        = "resum-api-lambda-policy"
  description = "IAM policy for Lambda to use dynamoDB"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = [
        "dynamodb:*",
        "iam:GetPolicy",
        "iam:GetPolicyVersion",
        "iam:GetRole",
        "iam:GetRolePolicy",
        "iam:ListAttachedRolePolicies",
        "iam:ListRolePolicies",
        "iam:ListRoles",
        "lambda:*",
        "logs:DescribeLogGroups",
        "states:DescribeStateMachine",
        "states:ListStateMachines",
        "tag:GetResources",
        "xray:GetTraceSummaries",
        "xray:BatchGetTraces"
      ],
      Effect   = "Allow",
      Resource = "*"
      },
      {
        "Effect" : "Allow",
        "Action" : "iam:PassRole",
        "Resource" : "*",
        "Condition" : {
          "StringEquals" : {
            "iam:PassedToService" : "lambda.amazonaws.com"
          }
        }
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "logs:DescribeLogStreams",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ],
        "Resource" : "arn:aws:logs:*:*:log-group:/aws/lambda/*"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "attach_email_lambda_policy" {
  policy_arn = aws_iam_policy.resum_api_policy.arn
  role       = aws_iam_role.resum_api_lambda_role.name
}
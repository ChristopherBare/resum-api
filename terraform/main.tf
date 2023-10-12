data "github_repository" "repo" {
  full_name = var.github_repo_full.name
}

locals {
  lambda_package_checksum = filebase64(var.lambda_package_path)
}

resource "random_id" "unique_id" {
  byte_length = 8 # You can adjust the length as needed
}

resource "aws_lambda_function" "resum_api_lambda" {
  function_name = "resum-api-lambda-${var.github_repo}-${var.branch_name}"
  role          = data.aws_iam_role.lambda_exec_role.arn
  handler       = "main"
  runtime       = "go1.21"
  source_code_hash = local.lambda_package_checksum
  filename = var.lambda_package_path
}

variable "lambda_package_path" {
  description = "Path to the Lambda deployment package"
  type        = string
  default     = "../lambda.zip"
}

data "aws_iam_role" "lambda_exec_role" {
  name = "lambda_role"
}



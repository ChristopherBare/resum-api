data "github_repository" "repo" {
  full_name = var.github_repo_full.name
}

resource "random_id" "unique_id" {
  byte_length = 8 # You can adjust the length as needed
}

resource "aws_lambda_function" "resum_api_lambda" {
  function_name = "resum-api-lambda-${var.github_repo}-${var.branch_name}"
  role          = data.aws_iam_role.lambda_exec_role.arn
  handler       = "main"
  runtime       = "go1.x"

  filename = "../zip/lambda.zip"
}

data "aws_iam_role" "lambda_exec_role" {
  name = "lambda_role"
}



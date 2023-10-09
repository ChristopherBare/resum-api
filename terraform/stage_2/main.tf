data "terraform_remote_state" "stage_1" {
  backend = "s3"

  config = {
    bucket = "terraform-backend-state-bucket-lambda-resum-api"
    key    = "terraform.tfstate" # Use the same key as in Stage 1
    region = "us-east-1"         # Use the same region as in Stage 1
  }
}

resource "random_id" "unique_id" {
  byte_length = 8 # You can adjust the length as needed
}

resource "aws_lambda_function" "resum_api_lambda" {
  function_name = "resum-api-lambda"
  role          = data.aws_iam_role.lambda_exec_role.arn
  handler       = "main"
  runtime       = "go1.x"
  s3_bucket     = data.terraform_remote_state.stage_1.outputs.bucket_name
  s3_key        = "lambda.zip"
  depends_on    = [data.terraform_remote_state.stage_1]
}

data "aws_iam_role" "lambda_exec_role" {
  name = "lambda_role"
}





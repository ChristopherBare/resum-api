data "terraform_remote_state" "stage_1" {
  backend = "s3"

  config = {
    bucket         = "terraform-backend-state-bucket-lambda-resum-api"
    key            = "terraform.tfstate"  # Use the same key as in Stage 1
    region         = "us-east-1"           # Use the same region as in Stage 1
  }
}

resource "aws_lambda_function" "resum_api_lambda" {
  function_name    = "resum-api-lambda"
  role             = aws_iam_role.lambda_exec_role.arn
  handler          = "main"
  runtime          = "go1.x"
  s3_bucket = data.terraform_remote_state.stage_1.outputs.bucket_name
  s3_key = "lambda.zip"
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_role-${timestamp()}"

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





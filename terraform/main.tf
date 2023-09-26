resource "aws_lambda_function" "resum_api_lambda" {
  filename         = "lambda.zip"
  function_name    = "resum-api-lambda"
  role             = aws_iam_role.lambda_exec_role.arn
  handler          = "main" # Change to your Go function's handler name
  runtime          = "go1.x"

  # Use the S3 bucket as the source code
  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = "lambda.zip"
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_role"

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





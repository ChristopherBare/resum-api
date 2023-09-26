output "lambda_function_arn" {
  value = aws_lambda_function.resum_api_lambda.arn
}

output "api_gateway_invoke_url" {
  value = aws_api_gateway_deployment.api_deployment.invoke_url
}

output "lambda_bucket_name" {
  value = aws_s3_bucket.lambda_bucket.bucket
}

output "lambda_function_name" {
  value = aws_lambda_function.resum_api_lambda.function_name
}
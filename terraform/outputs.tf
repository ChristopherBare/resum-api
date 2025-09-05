output "lambda_function_arn" {
  value = aws_lambda_function.resum_api_lambda.arn
}

output "api_gateway_invoke_url" {
  value = "https://${aws_api_gateway_rest_api.resum_api.id}.execute-api.${data.aws_region.current.id}.amazonaws.com/${aws_api_gateway_stage.prod.stage_name}"
}

output "lambda_function_name" {
  value = aws_lambda_function.resum_api_lambda.function_name
}
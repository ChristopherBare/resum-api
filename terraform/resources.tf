resource "aws_api_gateway_rest_api" "resum_api" {
  name        = "resum-api"
  description = "A resume, as an API"
}

resource "aws_api_gateway_resource" "api_resource" {
  parent_id = aws_api_gateway_rest_api.resum_api.root_resource_id
  path_part = "resume"
  rest_api_id = aws_api_gateway_rest_api.resum_api.id
}

resource "aws_api_gateway_method" "api_method" {
  rest_api_id   = aws_api_gateway_rest_api.resum_api.id
  resource_id   = aws_api_gateway_resource.api_resource.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "api_integration" {
  rest_api_id = aws_api_gateway_rest_api.resum_api.id
  resource_id = aws_api_gateway_resource.api_resource.id
  http_method = aws_api_gateway_method.api_method.http_method

  integration_http_method = "POST"
  type                   = "AWS_PROXY"
  uri                    = aws_lambda_function.resum_api_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "api_deployment" {
  depends_on = [aws_api_gateway_integration.api_integration]
  rest_api_id = aws_api_gateway_rest_api.resum_api.id
  stage_name  = "prod"
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.resum_api_lambda.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = aws_api_gateway_deployment.api_deployment.execution_arn
}

resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "lambda-bucket"
}




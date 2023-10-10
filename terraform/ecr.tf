resource "aws_ecr_repository" "lambda_repo" {
  name = "resum-api-lambda-repo"
}

resource "aws_ecr_lifecycle_policy" "my_lambda_repo_policy" {
  repository = aws_ecr_repository.lambda_repo.name

  policy = jsonencode({
    rules = [
      {
        rulePriority = 10,
        tagStatus    = "untagged",
        maxImageAge  = 7
      },
      {
        rulePriority = 20,
        tagStatus    = "any",
        countType    = "imageCountMoreThan",
        countNumber  = 5,
        countUnit    = "image"
      },
    ]
  })
}

resource "null_resource" "build_and_push" {
  triggers = {
    ecr_repository_id = aws_ecr_repository.lambda_repo.id
  }

  provisioner "local-exec" {
    command = <<EOT
    # Build and tag the Docker image
    docker build -t resum-api-lambda .

    # Tag the image with the ECR repository URI
    docker tag resum-api-lambda:latest ${aws_ecr_repository.lambda_repo.repository_url}:latest

    # Authenticate Docker to your ECR registry
    aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin ${aws_ecr_repository.lambda_repo.repository_url}

    # Push the Docker image to ECR
    docker push ${aws_ecr_repository.lambda_repo.repository_url}:latest
EOT
  }
}

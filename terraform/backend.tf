terraform {
  backend "s3" {
    bucket         = "terraform-backend-state-bucket-lambda-resum-api"
    key            = "terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
  }
}

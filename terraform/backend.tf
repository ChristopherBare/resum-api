terraform {
  backend "s3" {
    bucket = "terraform-backend-state-bucket-lambda-resum-api"
    key    = "terraform.tfstate"
    region = "us-east-1"
    dynamodb_table = "terraform-lock-table-resum-api"
  }
}

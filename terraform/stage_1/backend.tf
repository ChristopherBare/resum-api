#terraform {
#  backend "s3" {
#    bucket         = "terraform-backend-state-bucket-lambda-resum-api" # Replace with your S3 bucket name
#    key            = "terraform.tfstate"         # Replace with your state file name
#    region         = "us-east-1"                 # Replace with your desired AWS region
##    encrypt        = true                        # Optional: enable encryption
##    dynamodb_table = "my-lock-table"             # Optional: use a DynamoDB table for locking
#  }
#}

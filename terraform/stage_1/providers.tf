terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    github = {
      source  = "integrations/github"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  access_key = "AKIA2SCXB3FXHYRB2JVQ"
  secret_key = "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"
}
provider "github" {
}

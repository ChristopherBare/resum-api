terraform {
  required_providers {
    github = {
      source  = "integrations/github"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

provider "github" {}

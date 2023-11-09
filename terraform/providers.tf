terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

variable "GITHUB_TOKEN" {}

provider "github" {
  token = var.GITHUB_TOKEN != "" ? var.GITHUB_TOKEN : env("GITHUB_TOKEN")
}
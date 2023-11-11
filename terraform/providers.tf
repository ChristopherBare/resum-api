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

variable "GITHUB_TOKEN" {
  description = "Secret value for token"
}

locals {
  github_token_secret_value = "%s"
}

resource "github_actions_environment_secret" "GITHUB_TOKEN" {
  repository      = data.github_repository.repo.name
  environment     = "prod"
  secret_name     = "MY_GITHUB_TOKEN"
  plaintext_value = local.github_token_secret_value
}

provider "github" {
  token = local.github_token_secret_value
}

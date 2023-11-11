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

resource "github_repository_environment" "repo_environment" {
  repository   = data.github_repository.repo.name
  environment  = "prod"
}

resource "github_actions_environment_secret" "GITHUB_TOKEN" {
  repository        = data.github_repository.repo.name
  environment       = github_repository_environment.repo_environment.environment
  secret_name       = "GITHUB_TOKEN"
  plaintext_value   = var.GITHUB_TOKEN
}

provider "github" {
  token = github_actions_environment_secret.GITHUB_TOKEN.plaintext_value
}

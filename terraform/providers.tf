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

variable "GITHUB_TOKEN"{description = "secret value for token"}

resource "github_actions_environment_secret" "GITHUB_TOKEN" {
  repository = data.github_repository.repo
  environment       = github_repository_environment.repo_environment
  secret_name       = "GITHUB_TOKEN"
  plaintext_value   = var.GITHUB_TOKEN
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}
provider "github" {
  token = github_actions_environment_secret.GITHUB_TOKEN
}
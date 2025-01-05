terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {}
}

provider "aws" {
  region = var.aws_region
}

resource "aws_ecr_repository" "simplebank_repository" {
  name                 = var.ecr_repository_name
  image_tag_mutability = var.ecr_image_tag_mutability
  image_scanning_configuration {
    scan_on_push = var.ecr_scan_on_push
  }
  encryption_configuration {
    encryption_type = var.ecr_encryption_type
  }

  tags = {
    Environment = "Development"
    Project     = "MyProject"
  }
}

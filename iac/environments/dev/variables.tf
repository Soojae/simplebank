variable "ecr_repository_name" {
  description = "Name of the ECR repository"
  type        = string
}

variable "ecr_image_tag_mutability" {
  description = "Image tag mutability for the ECR repository"
  type        = string
}

variable "ecr_scan_on_push" {
  description = "Enable image scan on push for the ECR repository"
  type        = bool
}

variable "ecr_encryption_type" {
  description = "Encryption type for the ECR repository"
  type        = string
}

variable "aws_region" {
  description = "AWS region"
  type        = string
}

variable "db_identifier" {
  description = "Database identifier"
  type        = string
}

variable "db_engine_version" {
  description = "PostgreSQL engine version"
  type        = string
}

variable "db_instance_class" {
  description = "Instance class for RDS"
  type        = string
}

variable "db_allocated_storage" {
  description = "Allocated storage in GB"
  type        = number
}

variable "db_storage_type" {
  description = "Storage type"
  type        = string
}

variable "db_username" {
  description = "Master username for RDS"
  type        = string
}

variable "db_subnet_group_name" {
  description = "DB subnet group name"
  type        = string
}

variable "publicly_accessible" {
  description = "Whether the RDS instance is publicly accessible"
  type        = bool
}

variable "multi_az" {
  description = "Whether to deploy RDS in Multi-AZ"
  type        = bool
}

variable "storage_encrypted" {
  description = "Whether to encrypt the RDS instance"
  type        = bool
}

variable "existing_sg_name" {
  description = "Name of the existing Security Group"
  type        = string
}

variable "db_name" {
  description = "Database name"
  type        = string
}
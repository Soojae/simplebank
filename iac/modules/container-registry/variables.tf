variable "ecr_repository_name" {
  description = "Name of the ECR repository"
  type        = string
}

variable "ecr_image_tag_mutability" {
  description = "Image tag mutability for the ECR repository"
  type        = string
  default     = "MUTABLE"
}

variable "ecr_scan_on_push" {
  description = "Enable image scan on push for the ECR repository"
  type        = bool
  default     = true
}

variable "ecr_encryption_type" {
  description = "Encryption type for the ECR repository"
  type        = string
  default     = "AES256"
}

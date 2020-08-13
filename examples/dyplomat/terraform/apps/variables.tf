variable "dyplomat_image" {
  description = "Registry and path for your dyplomat image"
  default     = "myregistry/dyplomat:latest"
}

variable "dyplomat_role_arn" {
    description = "role ARN for dyplomat IAM role"
    default     = "arn:aws:iam::ACCOUNT_NO:role/dyplomat"
}

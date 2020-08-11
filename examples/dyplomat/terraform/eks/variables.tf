variable "region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "profile_name" {
  description = "AWS auth profile to use"
  default     = ""
}

variable "eks_subnets" {
  description = "List of CIDR subnet ids the eks cluster will be created with. Must be in 2 different AZs."
  default     = ["", ""]
}

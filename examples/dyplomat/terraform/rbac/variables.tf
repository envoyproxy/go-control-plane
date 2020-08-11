variable "oidc_provider_demo1" {
  description = "demo1-eks cluster oidc provider for EKS pod identity webhook"
  default     = "arn:aws:iam::ACCOUNT_NO:oidc-provider/oidc.eks.REGION.amazonaws.com/id/DEMO1_ID"
}

variable "oidc_provider_demo2" {
  description = "demo2-eks cluster oidc provider for EKS pod identity webhook"
  default     = "arn:aws:iam::ACCOUNT_NO:oidc-provider/oidc.eks.REGION.amazonaws.com/id/DEMO2_ID"
}

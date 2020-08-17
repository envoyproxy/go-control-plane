terraform {
  required_version = ">= 0.12"
  backend "local" {
    path = ""
  }
}
module "eks" {
  source = "./eks"
}

module "rbac" {
  source = "./rbac"
}

module "apps" {
  source = "./apps"
}

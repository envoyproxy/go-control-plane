provider "kubernetes" {
  version = "~> 1.11"
}

resource "kubernetes_namespace" "namespace" {
  metadata {
    name = "demo"
  }
}


resource "kubernetes_service_account" "dyplomat" {
  metadata {
    name = "dyplomat"
    namespace = "demo"
    annotations = {
      "eks.amazonaws.com/role-arn" = vars.dyplomat_role_arn
    }
  }
}

resource "kubernetes_cluster_role" "dyplomat" {
  metadata {
    name = "dyplomat"
  }

  rule {
    api_groups = [""]
    resources  = ["endpoints"]
    verbs      = ["get", "list", "watch"]
  }
}

resource "kubernetes_cluster_role_binding" "dyplomat" {
  metadata {
    name = "dyplomat"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "dyplomat"
  }

  subject {
    kind = "User"
    name = vars.dyplomat_role_arn
    api_group = "rbac.authorization.k8s.io"
    namespace = "demo"
  }
  subject {
    kind = "User"
    name = "system:serviceaccount:demo:dyplomat"
    api_group = "rbac.authorization.k8s.io"
    namespace = "demo"
  }
}

resource "kubernetes_deployment" "dyplomat" {
  metadata {
    name = "dyplomat"
    labels = {
      app = "dyplomat"
    }
    namespace = "demo"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "dyplomat"
      }
    }

    template {
      metadata {
        labels = {
          app = "dyplomat"
        }
      }

      spec {
        service_account_name = "dyplomat"
        automount_service_account_token = true

        container {
          image = vars.dyplomat_image
          name  = "dyplomat"
          image_pull_policy = "Always"

          port {
            container_port = "8080"
            protocol = "TCP"
            name = "http"
          }
        }
      }      
    }
  }
}

resource "kubernetes_service" "dyplomat" {
  metadata {
    name = "dyplomat"
    namespace = "demo"
  }
  spec {
    selector = {
      app = "dyplomat"
    }
    session_affinity = "ClientIP"
    port {
      port        = 8080
      target_port = 8080
    }
  }
}

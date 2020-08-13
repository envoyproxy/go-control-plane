resource "kubernetes_deployment" "frontenvoy" {
  metadata {
    name = "frontenvoy"
    labels = {
      app = "frontenvoy"
    }
    namespace = "demo"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "frontenvoy"
      }
    }

    template {
      metadata {
        labels = {
          app = "frontenvoy"
        }
      }

      spec {
        container {
          image = "envoyproxy/envoy:v1.14.4"
          name  = "frontenvoy"
          command = ["/usr/local/bin/envoy"]
          args = [
            "--config-path /etc/envoy/frontenvoy.yaml",
            "--service-cluster frontenvoy",
            "--service-node mesh",
          ]

          volume_mount {
            name = "envoy"
            mount_path = "/etc/envoy"
          }

          port {
            container_port = "8080"
            protocol = "TCP"
            name = "http"
          }
        }

        volume {
          name = "envoy"
          config_map {
            name = "frontenvoy-cm"
          }
        }
      }      
    }
  }
}

resource "kubernetes_config_map" "frontenvoy-cm" {
  metadata {
    name = "frontenvoy-cm"
    namespace = "demo"
  }

  data = {
    "frontenvoy.yaml" = "${file("${path.module}/frontenvoy.yaml")}"
  }
}

resource "kubernetes_service" "frontenvoy" {
  metadata {
    name = "frontenvoy"
    namespace = "demo"
  }
  spec {
    selector = {
      app = "frontenvoy"
    }
    session_affinity = "ClientIP"
    port {
      port        = 8080
      target_port = 8080
    }

    cluster_ip = "None"
  }
}

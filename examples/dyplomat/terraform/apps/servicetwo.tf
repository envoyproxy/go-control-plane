resource "kubernetes_deployment" "servicetwo" {
  metadata {
    name = "servicetwo"
    labels = {
      app = "servicetwo"
    }
    namespace = "demo"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "servicetwo"
      }
    }

    template {
      metadata {
        labels = {
          app = "servicetwo"
        }
      }

      spec {
        container {
          image = "envoyproxy/envoy:v1.14.4"
          name  = "envoysidecar"
          command = ["/usr/local/bin/envoy"]
          args = [
            "--config-path /etc/envoy/envoysidecar.yaml",
            "--mode serve",
            "--service-cluster servicetwo",
            "--service-node mesh"
          ]

          volume_mount {
            name = "envoy"
            mount_path = "/etc/envoy"
          }

          port {
            container_port = "8000"
            protocol = "TCP"
            name = "http"
          }

          port {
            container_port = "8001"
            protocol = "TCP"
            name = "envoy-admin"
          }
        }

        container {
          image = "nginx:latest"
          name  = "servicetwo"
        }

        volume {
          name = "envoy"
          config_map {
            name = "servicetwo-cm"
          }
        }
      }      
    }
  }
}

resource "kubernetes_config_map" "servicetwo-cm" {
  metadata {
    name = "servicetwo-cm"
    namespace = "demo"
  }

  data = {
    "envoysidecar.yaml" = "${file("${path.module}/envoysidecar.yaml")}"
  }
}

resource "kubernetes_service" "servicetwo" {
  metadata {
    name = "servicetwo"
    namespace = "demo"
    labels = {
      xds = "true"
    }
  }
  spec {
    selector = {
      app = "servicetwo"
    }
    session_affinity = "ClientIP"
    port {
      port        = 8000
      target_port = 8000
    }

    type = "NodePort"
  }
}

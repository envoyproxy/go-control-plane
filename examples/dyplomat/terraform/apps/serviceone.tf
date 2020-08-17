resource "kubernetes_deployment" "serviceone" {
  metadata {
    name = "serviceone"
    labels = {
      app = "serviceone"
    }
    namespace = "demo"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "serviceone"
      }
    }

    template {
      metadata {
        labels = {
          app = "serviceone"
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
            "--service-cluster serviceone",
            "--service-node mesh",
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
          name  = "serviceone"
        }

        volume {
          name = "envoy"
          config_map {
            name = "serviceone-cm"
          }
        }
      }      
    }
  }
}

resource "kubernetes_config_map" "serviceone-cm" {
  metadata {
    name = "serviceone-cm"
    namespace = "demo"
  }

  data = {
    "envoysidecar.yaml" = "${file("${path.module}/envoysidecar.yaml")}"
  }
}

resource "kubernetes_service" "serviceone" {
  metadata {
    name = "serviceone"
    namespace = "demo"
    labels = {
      xds = "true"
    }
  }
  spec {
    selector = {
      app = "serviceone"
    }
    session_affinity = "ClientIP"
    port {
      port        = 8000
      target_port = 8000
    }

    type = "NodePort"
  }
}

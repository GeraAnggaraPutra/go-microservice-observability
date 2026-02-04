logging {
  level = "info"
}

prometheus.receive_http "default" {
  forward_to = [prometheus.remote_write.to_prometheus.receiver]
}

prometheus.remote_write "to_prometheus" {
  endpoint {
    url = "http://shared-prometheus:9090/api/v1/write"
  }
}

loki.write "default" {
  endpoint {
    url = "http://shared-loki:3100/loki/api/v1/push"
  }
}

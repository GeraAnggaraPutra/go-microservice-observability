logging {
  level = "info"
}

// METRICS
prometheus.scrape "service_b" {
  targets = [
    {
      "__address__" = "api:8001",
      "job"         = "service-b",
      "service"     = "service-b",
    },
  ]
  forward_to = [prometheus.remote_write.to_central.receiver]
}

prometheus.remote_write "to_central" {
  endpoint {
    url = "http://shared-prometheus:9090/api/v1/write"
  }
}

// LOGS
discovery.docker "containers" {
  host = "unix:///var/run/docker.sock"
}

discovery.relabel "filter_service_b" {
  targets = discovery.docker.containers.targets

  rule {
    source_labels = ["__meta_docker_container_name"]
    regex         = "/service-b-api"
    action        = "keep"
  }
}

loki.source.docker "service_b_logs" {
  host       = "unix:///var/run/docker.sock"
  targets    = discovery.relabel.filter_service_b.output
  forward_to = [loki.process.filter_logs.receiver]
}

loki.process "filter_logs" {
  // Extract 'path' field from JSON log content
  stage.json {
    expressions = {
      path = "path",
    }
  }

  // Discard logs if path contains /metrics or /health
  stage.drop {
    source = "path"
    expression = "/metrics|/health"
  }

  // Add an identification label
  stage.static_labels {
    values = {
      "job"     = "service-b",
      "service" = "service-b",
    }
  }

  forward_to = [loki.write.to_central.receiver]
}

loki.write "to_central" {
  endpoint {
    url = "http://shared-loki:3100/loki/api/v1/push"
  }
}
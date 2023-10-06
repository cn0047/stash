service {
  name = "hwc1"
  id = "hwc-1"
  port = 8081

  connect {
    sidecar_service {
      proxy {
        upstreams {
          destination_name = "hws1"
          local_bind_port = 8060
        }
      }
    }
  }

  check {
    id       = "hwc1-check"
    http     = "http://localhost:8081/"
    method   = "GET"
    interval = "1s"
    timeout  = "1s"
  }
}

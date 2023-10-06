service {
  name = "hws1"
  id = "hws-1"
  port = 8080

  connect {
    sidecar_service {}
  }

  check {
    id       = "hws1-check"
    http     = "http://localhost:8080/"
    method   = "GET"
    interval = "1s"
    timeout  = "1s"
  }
}

services {
  id = "red0"
  name = "redis"
  tags = [
    "primary"
  ]
  address = ""
  port = 6000
  checks = [
    {
      args = ["/bin/check_redis", "-p", "6000"]
      interval = "5s"
      timeout = "20s"
    }
  ]
}
services {
  id = "red1"
  name = "redis"
  tags = [
    "delayed",
    "secondary"
  ]
  address = ""
  port = 7000
  checks = [
    {
      args = ["/bin/check_redis", "-p", "7000"]
      interval = "30s"
      timeout = "60s"
    }
  ]
}

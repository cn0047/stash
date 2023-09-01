service {
  name = "redis"
  id   = "redis"
  port = 80
  tags = ["primary"]

  meta = {
    custom_meta_key = "custom_meta_value"
  }

  tagged_addresses = {
    lan = {
      address = "192.168.0.55"
      port    = 8000
    }

    wan = {
      address = "198.18.0.23"
      port    = 80
    }
  }
}

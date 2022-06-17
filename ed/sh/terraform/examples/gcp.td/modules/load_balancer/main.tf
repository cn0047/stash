resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  provider = "google-beta"
  for_each = toset(var.regions)

  name                  = "cloudrun-neg"
  network_endpoint_type = "SERVERLESS"
  region                = each.value

  cloud_run {
    service = var.name
  }
}

resource "google_compute_backend_service" "default" {
  name     = "${var.name}-backend-service"
  protocol = "HTTPS"

  dynamic "backend" {
    for_each = toset(var.regions)
    content {
      group = "https://www.googleapis.com/compute/beta/projects/${var.project_id}/regions/${backend.value}/networkEndpointGroups/cloudrun-neg"
    }
  }

  log_config {
    enable      = true
    sample_rate = 1
  }

  depends_on = [
    google_compute_region_network_endpoint_group.cloudrun_neg
  ]
}

resource "google_compute_url_map" "default" {
  name            = var.name
  default_service = google_compute_backend_service.default.id
}

resource "random_id" "certificate" {
  byte_length = 4
  prefix      = "${var.name}-cert-"

  keepers = {
    domain = var.managed_domain
  }
}

resource "google_compute_managed_ssl_certificate" "default" {
  name = random_id.certificate.hex
  managed {
    domains = ["${var.managed_domain}."]
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_compute_target_https_proxy" "default" {
  name             = var.name
  url_map          = google_compute_url_map.default.id
  ssl_certificates = [google_compute_managed_ssl_certificate.default.id]
}

resource "google_compute_global_address" "default" {
  provider   = "google-beta"
  name         = var.name
}

resource "google_compute_global_forwarding_rule" "default" {
  provider   = "google-beta"
  project    = var.project_id
  name       = var.name
  target     = google_compute_target_https_proxy.default.id
  port_range = "443"
  ip_address = google_compute_global_address.default.id
}

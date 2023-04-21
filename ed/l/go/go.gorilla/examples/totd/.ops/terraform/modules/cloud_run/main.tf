resource "google_service_account" "default" {
  account_id   = var.service_account_name
  display_name = "Service account for the cloud run apps to run as"
}

resource "google_project_iam_member" "default" {
  for_each = toset(var.service_account_roles)

  project = var.project_id
  role    = each.value
  member  = "serviceAccount:${google_service_account.default.email}"
}

resource "google_cloud_run_service" "default" {
  provider = google-beta
  for_each = toset(var.regions)

  name     = var.service_name
  location = each.value

  metadata {
    annotations = {
     "run.googleapis.com/ingress" = "internal"
    }
  }

  template {
    spec {
      containers {
        image = var.docker_image
        env {
          name = "ENV"
          value = var.env_type
        }
        env {
          name = "HOST"
          value = "0.0.0.0"
        }
        env {
          name = "SPANNER_DATABASE"
          value = var.db_connection_string
        }
        env {
          name = "DB_LOG_ENABLED"
          value = var.enable_app_logs
        }
        env {
          name = "MFC_CONFIG_AUTO_REFRESH_TIME"
          value = var.mfc_config_auto_refresh_time
        }
      }
      service_account_name = google_service_account.default.email
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = all
  }
}

output "regional_urls" {
  value = tomap({
    for key, service in resource.google_cloud_run_service.default : key => service.status[0].url
  })
  description = "This is a map of region to cloud run service url deployed to that region."
}

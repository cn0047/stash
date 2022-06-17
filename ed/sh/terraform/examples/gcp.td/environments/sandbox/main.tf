terraform {
  required_version = ">= 1.0.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.12.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = ">= 4.12.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.provider_region
}

provider "google-beta" {
  project = var.project_id
  region  = var.provider_region
}

resource "google_project_service" "default" {
  for_each           = toset(var.google_apis)
  provider           = google
  service            = each.value
  disable_on_destroy = false
}

resource "google_project_iam_member" "dev" {
  for_each = toset(var.team_roles)
  project  = var.project_id
  role     = each.value
  member   = "group:${var.team_email}"
}

module "cloud_spanner" {
  spanner_instance_config = var.spanner_instance_config
  source                  = "../modules/cloud_spanner"
  depends_on = [
    google_project_service.default
  ]
}

module "cloud_run" {
  source     = "../modules/cloud_run"
  project_id = var.project_id
  env_type   = var.env_type
  regions    = var.regions

  docker_image = var.docker_image
  db_connection_string = var.db_connection_string

  depends_on = [
    google_project_service.default,
    module.cloud_spanner
  ]
}

module "load_balancer" {
  source         = "../modules/load_balancer"
  project_id     = var.project_id
  managed_domain = var.managed_domain
  managed_zone   = var.managed_zone
  regions        = var.regions
  env_type       = var.env_type
  depends_on = [
    google_project_service.default,
  ]
}

variable "project_id" {
  type = string
}

variable "project_name" {
  type = string
}

variable "env_type" {
  type = string
}

variable "provider_region" {
  type    = string
  default = "us-central1"
}

variable "regions" {
  type = list(string)
}

variable "spanner_instance_config" {
  type = string
}

variable "docker_image" {
  type = string
}

variable "db_connection_string" {
  type = string
}

variable "managed_domain" {
  type = string
}

variable "managed_zone" {
  type = string
}

variable "google_apis" {
  type = list(string)
  default = [
    // General
    "iam.googleapis.com",

    // DNS
    "dns.googleapis.com",

    // API Gateway
    "apigateway.googleapis.com",
    "servicemanagement.googleapis.com",
    "servicecontrol.googleapis.com",

    // Spanner
    "spanner.googleapis.com",

    // Cloud Tracing
    "cloudtrace.googleapis.com",

    // Cloud Run related
    "cloudresourcemanager.googleapis.com",
    "cloudbuild.googleapis.com",

    // Cloud Run
    "artifactregistry.googleapis.com",
    "run.googleapis.com",
  ]
}

variable "team_email" {
  type    = string
}

variable "team_roles" {
  type = list(string)
  default = [
    "roles/spanner.databaseUser",
    "roles/cloudfunctions.invoker",
    "roles/compute.loadBalancerAdmin"
  ]
}

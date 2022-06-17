variable "project_id" {
  type = string
}

variable "service_account_name" {
  type    = string
  default = "cloud-run-sa"
}

variable "service_account_roles" {
  type = list(string)
  default = [
    "roles/spanner.databaseUser",
    "roles/logging.logWriter",
    "roles/cloudtrace.agent",
    "roles/monitoring.metricWriter"
  ]
}

variable "regions" {
  type = list(string)
}

variable "service_name" {
  type    = string
}

variable "env_type" {
  type = string
}

variable "docker_image" {
  type    = string
}

variable "db_connection_string" {
  type    = string
}

variable "enable_app_logs" {
  type    = bool
  default = true
}

variable "mfc_config_auto_refresh_time" {
  type    = number
  default = 300
}

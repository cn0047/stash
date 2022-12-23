variable "project_id" {
  type = string
}

variable "spanner_instance_name" {
  type = string
}

variable "spanner_instance_display_name" {
  type    = string
  default = "Cloud Spanner"
}

variable "spanner_instance_config" {
  type = string
}

variable "spanner_instance_processing_units" {
  type    = number
  default = 100
}

variable "spanner_database_names" {
  type = list(string)
}

variable "spanner_backup_expiration_sec" {
  type = number
}

variable "spanner_backup_workflow_region" {
  type = string
}

variable "spanner_backup_service_account_id" {
  type = string
}

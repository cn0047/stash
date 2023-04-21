variable "spanner_instance_name" {
  type    = string
  default = "td"
}

variable "spanner_instance_config" {
  type = string
}

variable "spanner_instance_processing_units" {
  type    = number
  default = 100
}

variable "spanner_database_name" {
  type    = string
  default = "td"
}

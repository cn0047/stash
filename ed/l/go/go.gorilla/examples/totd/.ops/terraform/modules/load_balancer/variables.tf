variable "name" {
  type    = string
  default = "td"
}

variable "env_type" {
  type = string
}

variable "regions" {
  type = list(string)
}

variable "project_id" {
  type = string
}

variable "managed_zone" {
  type = string
}

variable "managed_domain" {
  type = string
}

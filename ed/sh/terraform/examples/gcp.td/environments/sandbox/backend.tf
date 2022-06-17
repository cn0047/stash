terraform {
  backend "gcs" {
    bucket = "sandbox-bkt"
    prefix = "terraform-state"
  }
}

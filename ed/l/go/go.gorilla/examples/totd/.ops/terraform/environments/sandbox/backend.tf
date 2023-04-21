terraform {
  backend "gcs" {
    bucket = "sandbox-20211128-sy7ccu-td-sandbox"
    prefix = "terraform-state"
  }
}

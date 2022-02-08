provider "google" {
  project = ""
}

resource "google_storage_bucket" "tf_storage_test_1" {
  name          = "tf_storage_test_1"
  location      = "EU"
  force_destroy = true
}

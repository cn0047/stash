provider "google" {
  project = ""
  region  = "us-central1"
}

resource "google_storage_bucket" "bucket" {
  name          = "tf_func_storage_1"
  location      = "EU"
}

resource "google_storage_bucket_object" "archive" {
  name   = "func.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./func.zip"
}

resource "google_cloudfunctions_function" "function" {
  name        = "func-test"
  runtime     = "go116"
  entry_point = "MainFunc"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = true
}

resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.function.project
  cloud_function = google_cloudfunctions_function.function.name
  region         = "us-central1"

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

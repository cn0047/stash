resource "google_spanner_instance" "dbi" {
  name             = var.spanner_instance_name
  display_name     = var.spanner_instance_name
  config           = var.spanner_instance_config
  processing_units = var.spanner_instance_processing_units
}

resource "google_spanner_database" "db" {
  instance = google_spanner_instance.dbi.name
  name     = var.spanner_database_name
}

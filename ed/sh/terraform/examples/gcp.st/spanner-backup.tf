resource "google_workflows_workflow" "backup_workflow" {
  for_each = google_spanner_database.default

  name            = "backup-${each.value.name}"
  region          = var.spanner_backup_workflow_region
  description     = "Spanner Backup Workflow - ${each.value.name}"
  service_account = google_service_account.backup_sa.id
  depends_on = [
    google_spanner_database.default
  ]
  source_contents = <<-EOF
  - first_step:
      assign:
      - expireTime: $${sys.now() + ${var.spanner_backup_expiration_sec}}
      - backupId: $${"${each.value.name}-" + text.replace_all(string(sys.now() * 1000), ".", "_")}
      - formattedExpireTime: $${time.format(expireTime)}
  - createBackup:
      call: http.post
      args:
          url: $${"https://content-spanner.googleapis.com/v1/projects/${var.project_id}/instances/${google_spanner_instance.default.name}/backups?backupId=" + backupId}
          body:
              expireTime: $${formattedExpireTime}
              database: projects/${var.project_id}/instances/${google_spanner_instance.default.name}/databases/${each.value.name}
          auth:
              type: OAuth2
      result: createBackupResult
  - returnOutput:
          return: $${createBackupResult.code}
EOF
}

resource "google_cloud_scheduler_job" "backup_spanner" {
  for_each    = resource.google_workflows_workflow.backup_workflow
  name        = each.value.name
  description = "${each.value.name} Backup job"
  project     = var.project_id
  schedule    = "0 0 * * *"
  region      = var.spanner_backup_workflow_region

  http_target {
    http_method = "POST"
    uri         = "https://workflowexecutions.googleapis.com/v1/${each.value.id}/executions"

    oauth_token {
      service_account_email = google_service_account.backup_sa.email
    }
  }
}

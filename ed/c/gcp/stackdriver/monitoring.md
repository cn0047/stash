Monitoring
-

[api](https://cloud.google.com/monitoring/api/v3)

#### Alerting

Alerting policy sends event to notification channels.

Notification channel:
* Mobile device (via Cloud Mobile App).
* PagerDuty service.
* PagerDuty sync.
* Slack.
* Webhooks.
* Email.
* SMS.
* Pub/Sub.

Webhook payload examples:

````json
{
  "version": "test",
  "incident": {
    "incident_id": "12345",
    "scoping_project_id": "12345",
    "scoping_project_number": 12345,
    "url": "http://www.example.com",
    "started_at": 0,
    "ended_at": 0,
    "state": "OPEN",
    "summary": "Test Incident",
    "apigee_url": "http://www.example.com",
    "observed_value": "1.0",
    "resource": {
      "type": "example_resource",
      "labels": {
        "example": "label"
      }
    },
    "resource_type_display_name": "Example Resource Type",
    "resource_id": "12345",
    "resource_display_name": "Example Resource",
    "resource_name": "projects/12345/example_resources/12345",
    "metric": {
      "type": "test.googleapis.com/metric",
      "displayName": "Test Metric",
      "labels": {
        "example": "label"
      }
    },
    "metadata": {
      "system_labels": {
        "example": "label"
      },
      "user_labels": {
        "example": "label"
      }
    },
    "policy_name": "projects/12345/alertPolicies/12345",
    "policy_user_labels": {
      "example": "label"
    },
    "documentation": "Test documentation",
    "condition": {
      "name": "projects/12345/alertPolicies/12345/conditions/12345",
      "displayName": "Example condition",
      "conditionThreshold": {
        "filter": "metric.type=\"test.googleapis.com/metric\" resource.type=\"example_resource\"",
        "comparison": "COMPARISON_GT",
        "thresholdValue": 0.5,
        "duration": "0s",
        "trigger": {
          "count": 1
        }
      }
    },
    "condition_name": "Example condition",
    "threshold_value": "0.5"
  }
}

{
  "incident": {
    "condition": {
      "conditionThreshold": {
        "aggregations": [{"alignmentPeriod": "60s","perSeriesAligner": "ALIGN_MEAN"}],
        "comparison": "COMPARISON_GT",
        "duration": "0s",
        "filter": "resource.type = \"cloud_function\" AND resource.labels.function_name = \"hw\" AND metric.type = \"logging.googleapis.com/log_entry_count\" AND metric.labels.severity = \"ERROR\"",
        "trigger": {
          "count": 1
        }
      },
      "displayName": "Cloud Function - Log entries",
      "name": "projects/sandbox-2022-x1/alertPolicies/1130251226664066904/conditions/1130251226664063323"
    },
    "condition_name": "Cloud Function - Log entries",
    "documentation": {
      "content": "some docs here...",
      "mime_type": "text/markdown"
    },
    "ended_at": null,
    "incident_id": "0.mp9ojt5e6800",
    "metadata": {
      "system_labels": {},
      "user_labels": {}
    },
    "metric": {
      "displayName": "Log entries",
      "labels": {
        "log": "cloudfunctions.googleapis.com/cloud-functions"
      },
      "type": "logging.googleapis.com/log_entry_count"
    },
    "observed_value": "1.000",
    "policy_name": "1stTest",
    "policy_user_labels": {
      "foo": "bar"
    },
    "resource": {
      "labels": {
        "function_name": "hw",
        "project_id": "sandbox-2022-x1",
        "region": "us-central1"
      },
      "type": "cloud_function"
    },
    "resource_id": "",
    "resource_name": "sandbox-2022-x1 Cloud Function labels {project_id=sandbox-2022-x1, function_name=hw, region=us-central1}",
    "resource_type_display_name": "Cloud Function",
    "scoping_project_id": "sandbox-2022-x1",
    "scoping_project_number": 292209335709,
    "started_at": 1668792118,
    "state": "open",
    "summary": "Log entries for sandbox-2022-x1 Cloud Function labels {project_id=sandbox-2022-x1, function_name=hw, region=us-central1} with metric labels {log=cloudfunctions.googleapis.com/cloud-functions} is above the threshold of 0.000 with a value of 1.000.",
    "threshold_value": "0",
    "url": "https://console.cloud.google.com/monitoring/alerting/incidents/0.mp9ojt5e6800?project=sandbox-2022-x1"
  },
  "version": "1.2"
}
````

Scheduler
-

[docs](https://cloud.google.com/scheduler/docs)
[REST API](https://cloud.google.com/scheduler/docs/reference/rest)

Cloud Scheduler is per region.
Cloud Scheduler type: HTTP, PubSub, AppEngine HTTP, Workflows.

HTTP scheduler can specify headers as well.

Retry config:
* Max retry attempts (count).
* Max retry duration (time limit for retrying failed job).
* Min backoff duration (min time to wait before retrying).
* Max backoff duration (max time to wait before retrying).

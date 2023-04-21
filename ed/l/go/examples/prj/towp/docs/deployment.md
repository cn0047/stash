Deployment
-

To deploy wp please use github actions.

To deploy CloudRun please use `workflow_dispatch` with "CD" action,
<br>and provide: `branch`, `environment` and `git tag` (optional).
<br>This action will migrate DB as well for nonprod environments.

To migrate Cloud Spanner please use `workflow_dispatch` with "Migrate DB schema" action,
<br>and provide: `branch` and `environment`.

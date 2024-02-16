# helm examples

d=ed/sh/helm/examples/cronjob
helm dependency update $d/spanner-drop-testing-dbs
#
helm upgrade --debug spanner-drop-testing-dbs \
$d/spanner-drop-testing-dbs \
-f $d/values/all_common.yaml \
-f $d/values/testing.yaml \
--namespace="no-testing" \
--set-string image="us.gcr.io/no-images/go/no:master" \
--set-string global.no.project_id="no-testing" \
--set-string cmd_project="no-testing" \
--install \
--dry-run > /tmp/log.sh

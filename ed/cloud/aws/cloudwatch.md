CloudWatch
-

Metric > EC2 > NetworkIn - The number of bytes received on all network interfaces by the instance.

````sh
sudo service awslogs stop
````
````sh
aws --profile=$p cloudwatch list-dashboards
aws --profile=$p cloudwatch list-metrics | jq '.Metrics[].Namespace' | sort | uniq

aws cloudwatch put-metric-data --namespace 'prod.supervisor' --metric-name 'instance1.document' --value 1
aws cloudwatch put-metric-data --namespace 'prod.lf' --metric-name 'memoryfree' --unit Megabytes --value 9
````

Logs:

````sh
aws logs put-log-events --log-group-name cli_prod --log-stream-name x --log-events timestamp=`date +%s`,message=000

ln='/ecs/legacyfiles' # log name
st=1553270791 # start time
st=`date +%s`
p='cn911v2' # filter pattern
# get logs
aws logs filter-log-events \
    --log-group-name $ln \
    --start-time $st \
    --filter-pattern $p \
    --output json | jq '.events'
````

Logs insights:

````sh
filter @message like /cn911w2/
| fields @timestamp, @message
| sort @timestamp desc
| limit 20

stats count(*)
````

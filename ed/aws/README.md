AWS
-

[console](https://console.aws.amazon.com)
[go examples](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code)

````sh
# ubuntu
sudo apt-get install awscli

# osx
brew install awscli

aws configure
aws configure list

~/.aws/config
~/.aws/credentials

export AWS_REGION="eu-central-1"
export AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY=""
export AWS_SECRET_ACCESS_KEY=""
# or
export AWS_SECRET_KEY=""
````

````sh
--output=table|json|text
--debug # add this flag to get more debug info about command

--region us-east-1
# @see: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
us-east-1    # US East (N. Virginia)
eu-central-1 # EU (Frankfurt)
````

# CloudWatch

Metric > EC2 > NetworkIn - The number of bytes received on all network interfaces by the instance.

````sh
sudo service awslogs stop
````
````sh
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

# CodeDeploy

# CloudFront (CDN)

````sh
aws cloudfront list-distributions

aws cloudfront get-distribution --id $id
````

# Elastic Beanstalk

````
echo "web: application" > Procfile
````

# Route 53

````sh
aws route53 list-hosted-zones

aws route53 change-resource-record-sets --hosted-zone-id X1HGFN9JYF0T6U --change-batch '{
  "Comment": "s1",
  "Changes": [
    {
      "Action": "UPSERT",
      "ResourceRecordSet": {
        "Name": "x.net.com",
        "Type": "A",
        "TTL": 300,
        "ResourceRecords": [{"Value": "127.0.0.1"}]
      }
    }
  ]
}'
````

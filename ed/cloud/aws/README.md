AWS
-

[console](https://console.aws.amazon.com)
[quotas](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html)
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

--profile=x
--color=on
--output=table|json|text
--debug # add this flag to get more debug info about command

--region us-east-1
# @see: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html
us-east-1    # US East (N. Virginia)
eu-central-1 # EU (Frankfurt)
````

````sh
aws sts get-caller-identity
````

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

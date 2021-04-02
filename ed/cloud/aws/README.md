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
aws configure get aws_access_key_id
aws configure get aws_secret_access_key
# set configuration without prompts
aws configure set aws_access_key_id $k
aws configure set aws_secret_access_key $s

~/.aws/config
~/.aws/credentials

export AWS_REGION="eu-central-1"
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_ACCESS_KEY=""
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

Bastion host - special-purpose computer on a network specifically designed and configured to withstand attacks.

# AWS Cloud Map

AWS Cloud Map - service discovery.

# Elastic Beanstalk

EB - service for deploying and scaling web applications and services.

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

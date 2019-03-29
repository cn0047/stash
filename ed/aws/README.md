AWS
-

````sh
# ubuntu
sudo apt-get install awscli

# osx
brew install awscli

aws configure

~/.aws/config
````

````sh
--output=table|json|text
# add this flag to get more debug info about command
--debug
````

# S3 (Simple Storage Service)

````sh
# upload picture to s3
aws s3 cp /home/kovpak/Downloads/images.jpg s3://w3.stage.ziipr.bucket/test/x.jpg

aws s3 ls s3://bucket/img.png

# size of bucket and count of elements in bucket
aws s3api list-objects --bucket $bkt --query "[sum(Contents[].Size), length(Contents[])]"
````

# CloudWatch

````
sudo service awslogs stop
````
````sh
aws cloudwatch put-metric-data --namespace 'prod.supervisor' --metric-name 'instance1.document' --value 1
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

# SNS (Simple Notification Service)

Use to push message.

# SQS (Simple Queue Service)

Use to pull message.

# CodeDeploy

# CloudFront (CDN)

````sh
aws cloudfront list-distributions

aws cloudfront get-distribution --id $id
````

# (RDS) Relational Database Service

# (ECS) Elastic Container Service

[fargate pricing](https://aws.amazon.com/fargate/pricing/)

````sh
aws ecs list-clusters

aws ecs list-task-definitions

aws ecs list-services --cluster=clstr

# stop task, so AWS will create new one
aws ecs list-tasks --cluster=clstr --service=srvc
tid=`aws ecs list-tasks --cluster=clstr --service=srvc \
  | jq -r '.taskArns[0]' | awk -F '/' '{print $2}'`
aws ecs describe-tasks --cluster=clstr --task=$tid
aws ecs stop-task --cluster=clstr --task=$tid

# up service, so AWS will recreate tasks
aws ecs register-task-definition --cli-input-json file://ecs.taskDefinition.json
aws ecs update-service --cluster=clstr --service=srvc --task-definition=$tdid

aws ecs list-task-definitions --sort=DESC
aws ecs deregister-task-definition --task-definition="td:2"
````

# (ECR) Elastic Container Registry

````sh
aws ecr get-login --region us-east-1 --no-include-email
aws ecr create-repository --region us-east-1 --repository-name rName

aws ecr describe-images --region us-east-1 --repository-name legacy-files-dev
````

# Elastic Beanstalk

````
echo "web: application" > Procfile
````

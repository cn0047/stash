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
# add this flag to get more debug info about command
--debug
````

# S3 (Simple Storage Service)

````sh
# upload picture to s3
aws s3 cp /home/kovpak/Downloads/images.jpg s3://w3.stage.ziipr.bucket/test/x.jpg

aws s3 ls s3://bucket/img.png

# size of bucket and count of elements in bucket
aws s3api list-objects --bucket w3.stage.ziipr.bucket --query "[sum(Contents[].Size), length(Contents[])]"
````

# CloudWatch

````
sudo service awslogs stop
````
````
aws cloudwatch put-metric-data --namespace 'prod.supervisor' --metric-name 'instance1.document' --value 1

aws logs put-log-events --log-group-name cli_prod --log-stream-name x --log-events timestamp=`date +%s`,message=000
````

#### SNS (Simple Notification Service)

# SQS (Simple Queue Service)

# CodeDeploy

# (RDS) Relational Database Service

# (ECS) Elastic Container Service

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
````

# (ECR) Elastic Container Registry

````sh
aws ecr get-login --region us-east-1 --no-include-email
aws ecr create-repository --region us-east-1 --repository-name rName

aws ecr describe-images --region us-east-1 --repository-name legacy-files-dev
````

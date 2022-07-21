ECS (Elastic Container Service)
-

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

Makefile:
````sh
build_and_push_dev:
  docker build -t 107889011321.dkr.ecr.us-east-1.amazonaws.com/lf:dev \
    --build-arg APP_ENV=dev -f deploy/Dockerfile .
  eval $(shell aws ecr get-login --no-include-email --region us-east-1)
  docker push 107889011321.dkr.ecr.us-east-1.amazonaws.com/lf:dev

deploy_dev:
  aws ecs update-service --region=us-east-1 --cluster=prod-services --service=lf-dev --task-definition=`\
    aws ecs register-task-definition --region=us-east-1 --cli-input-json file://deploy/AWSTaskDefinition.dev.json \
      | grep --color=never -Eo 'lf-dev:[0-9]+' \
  `
  aws ecs list-task-definitions --sort=DESC \
    | grep --color=never -Eo 'lf-dev:[0-9]+' \
    | tail -n +6 \
    | while read v; do aws ecs deregister-task-definition --task-definition=$$v; done
````

# (ECR) Elastic Container Registry

````sh
aws ecr get-login --region us-east-1 --no-include-email
aws ecr create-repository --region us-east-1 --repository-name rName

aws ecr list-images --repository-name=x-eval

aws ecr describe-images --region us-east-1 --repository-name lf-dev
````

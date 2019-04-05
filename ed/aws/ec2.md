EC2 (Elastic Compute Cloud)
-

Default OS is Amazon Linux.

[instance types](https://aws.amazon.com/ec2/instance-types/)
[pricing](https://aws.amazon.com/ec2/pricing/)

Users (for ssh):
* ec2-user.
* ubuntu for AMI with ubuntu.

Launch instance:

On step 3: At `Advanced Details` in `User data` it is possible to write bash, like:
````sh
#!/bin/sh
yum -y install vim htop
````

From cli:

````sh
aws ec2 run-instances \
  --image-id ami-17a3e164 \
  --instance-type t2.medium \
  --count 1 \
  --security-groups ssh Ciklum web default \
  --key-name ziipr \
  --user-data 'echo 200' \
  --client-token KovpakTest4 \

  --instance-type t2.micro \

  --iam-instance-profile Name='aws-opsworks-ec2-role' \

  --user-data file://my_script.txt

# create new instance
aws ec2 run-instances \
  --image-id ami-0ae8c6f6de834bfca \
  --instance-type t2.large \
  --count 1 \
  --subnet-id subnet-2eadd305 \
  --iam-instance-profile Name='Kovpak-EC2' \
  --associate-public-ip-address \
  --key-name claws \
  --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=legacy-files-prod-3x}]' \
  --output json \
  > ec2.json
iId=`cat ec2.json | jq '.Instances[0].InstanceId'`
iId=`cat ec2.json | grep InstanceId --color=none | cut -f4 -d'"'`
# public dns name
h=`aws ec2 describe-instances --instance-ids $iId --query 'Reservations[].Instances[].PublicDnsName'`

# and add security groups to new instance
aws ec2 modify-instance-attribute \
  --instance-id $iId \
  --groups sg-bc0d83db sg-af58a3d7 sg-09cebaaf7cd6c61a6

aws ec2 describe-instance-status --instance-ids $iId | jq '.InstanceStatuses[0].InstanceState.Name'

# add new instance to lb
aws elbv2 register-targets \
  --target-group-arn $arn \
  --targets Id=$iId,Port=8080

````

Convenient information about instances:

````sh
aws ec2 describe-instances \
  --output table \
  --query 'Reservations[*].Instances[*].[Tags[0].Value,PublicDnsName,ImageId,LaunchTime,State.Name]' \
  --filter Name=tag:Name,Values=*prod*web*

  --filter Name=image-id,Values=ami-17a3e164
  --filter Name=dns-name,Values=ec2-52-51-65-182.eu-west-1.compute.amazonaws.com

````

One php session storage per several instances:

````
load balancer -> description -> port configuration = Stickiness: LBCookieStickinessPolicy, expirationPeriod='1800'
````

Auto Scaling:

Min     - AWS ensures that your group never goes below this size.
Max     - AWS ensures that your group never goes above this size.
Desired - AWS ensures that your group has this many instances.

Health check grace period - time to not health-check instance during launching instance.

#### AMI (Amazon Machine Image)

#### EBS (Elastic Block Store)

EBS provides persistent block storage volumes for use with Amazon EC2 instances in the AWS Cloud.

#### VPC (Virtual Private Cloud)

#### LB (Load Balancing)

* ELB (Elastic Load Balancing)
* ALB (Application Load Balancing)

#### ASG (Auto Scaling Group)

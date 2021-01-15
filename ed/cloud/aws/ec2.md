EC2 (Elastic Compute Cloud)
-

[instance types](https://aws.amazon.com/ec2/instance-types/)
[pricing](https://aws.amazon.com/ec2/pricing/)

Default OS is Amazon Linux.

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
# convenient information about instances:
aws --profile=$p ec2 describe-instances \
  --output table \
  --query 'Reservations[*].Instances[*].[Tags[0].Value,PublicDnsName,ImageId,LaunchTime,State.Name]' \

  --filter Name=tag:Name,Values=*prod*web*
  --filter Name=image-id,Values=ami-17a3e164
  --filter Name=dns-name,Values=ec2-52-51-65-182.eu-west-1.compute.amazonaws.com

# subnets
aws --profile=$p ec2 describe-subnets

# security groups
aws --profile=$p ec2 describe-security-groups

# list ssh keys
aws --profile=$p ec2 describe-key-pairs

# import ssh key
aws --profile=$p ec2 import-key-pair --key-name "defaultsshkey" \
  --public-key-material file://~/web/kovpak/gh/ed/sh/ssh/examples/nopwd/id_rsa.pub

# list vpcs
aws --profile=$p ec2 describe-vpcs

# create new ec2
ami=ami-00aa4671cbf840d82 # default Amazon Linux 2 AMI
aws ec2 run-instances \
  --image-id $ami \
  --instance-type t2.micro \
  --count 1 \
  --security-groups ssh cklm web default \
  --key-name zkey \
  --user-data 'echo 200' \
  --client-token KovpakTest4 \

  --instance-type t2.micro \

  --iam-instance-profile Name='aws-opsworks-ec2-role' \

  --user-data file://my_script.txt

# terminate
aws ec2 terminate-instances --instance-ids $(iId)

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
# public ip
ip=`aws ec2 describe-instances --instance-ids $iId --query 'Reservations[].Instances[].PublicIpAddress'`

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

ELB (Elastic Load Balancing) for multiple targets: EC2, containers, lambda funcs, etc.:
* ALB (Application Load Balancing) (Layer 7).
* Network Load Balancer (Layer 4) for TCP, UDP, TLS.
* Classic Load Balancer - basic LB for EC2.

#### ASG (Auto Scaling Group)

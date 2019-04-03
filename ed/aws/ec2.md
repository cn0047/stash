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

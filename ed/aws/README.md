AWS
-

## EC2 (Elastic Compute Cloud)

Launch instance:

On step 3: At `Advanced Details` in `User data` it is possible to write bash, like:
````sh
#!/bin/sh
yum -y install vim htop
````

From cli:

````
aws ec2 run-instances \
    --image-id ami-f9dd458a \
    --instance-type t2.micro \
    --count 1 \
    --security-groups ssh Ciklum web default \
    --key-name ziipr \
    --user-data IyEvYmluL3NoCnl1bSAteSBpbnN0YWxsIHZpbSBodG9wCg== \
    --client-token KovpakTest2 \

--user-data file://my_script.txt

````

Convenient information about instances:

````
aws ec2 describe-instances \
--output table \
--query 'Reservations[*].Instances[*].[Tags[0].Value,PublicDnsName,ImageId,State.Name]' \
--filter Name=tag:Name,Values=*prod*web*

--filter Name=dns-name,Values=ec2-52-51-65-182.eu-west-1.compute.amazonaws.com

````

One php session storage per several instances:

````
load balancer -> description -> port configuration = Stickiness: LBCookieStickinessPolicy, expirationPeriod='1800'
````

### CloudWatch

````
aws cloudwatch put-metric-data --namespace 'prod.supervisor' --metric-name 'instance1.document' --value 1
````

### AMI (Amazon Machine Image)

### EBS (Elastic Block Store)

### VPC (Virtual Private Cloud)

### ELB (Elastic Load Balancing)


## S3 (Simple Storage Service)

````
# s3
aws s3 cp /home/kovpak/Downloads/images.jpg s3://w3.stage.ziipr.bucket/test/x.jpg

# size of bucket and count of elements in bucket
aws s3api list-objects --bucket w3.stage.ziipr.bucket --query "[sum(Contents[].Size), length(Contents[])]"

````

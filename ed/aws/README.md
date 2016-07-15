AWS
-

#### CLI

#### ec2

````
# instances
aws ec2 describe-instances \
--output table \
--query 'Reservations[*].Instances[*].[Tags[0].Value,PublicDnsName,ImageId,State.Name]' \
--filter Name=tag:Name,Values=*prod*web*

````

#### s3

````
# s3
aws s3 cp /home/kovpak/Downloads/images.jpg s3://w3.stage.ziipr.bucket/test/x.jpg

# size of bucket and count of elements in bucket
aws s3api list-objects --bucket w3.stage.ziipr.bucket --query "[sum(Contents[].Size), length(Contents[])]"

````

#### One php session storage per several instances

````
load balancer -> description -> port configuration = Stickiness: LBCookieStickinessPolicy, expirationPeriod='1800'
````

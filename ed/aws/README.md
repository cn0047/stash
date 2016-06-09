AWS
-

#### cli

````
# instances
aws ec2 describe-instances \
--output table \
--query 'Reservations[*].Instances[*].[Tags[0].Value,PublicDnsName,ImageId,State.Name]' \
--filter Name=tag:Name,Values=*prod*web*

````

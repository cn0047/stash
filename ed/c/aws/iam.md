IAM (Identity and Access Management)
-

````sh
aws --profile=$p iam list-ssh-public-keys
aws --profile=$p iam list-access-keys

# list roles
aws --profile=$p iam list-roles | jq '.Roles[].RoleName'

# info about role
aws --profile=$p iam get-role --role-name=k8s
aws --profile=$p iam list-role-policies --role-name=k8s
aws --profile=$p iam list-attached-role-policies --role-name=k8s

#
aws --profile=$p iam create-role --role-name=k8s \
  --assume-role-policy-document file://~/iam.basic.json
````

````sh
# create group
aws iam --profile=$p create-group --group-name gk8s
aws iam --profile=$p attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess --group-name gk8s
aws iam --profile=$p attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess --group-name gk8s
aws iam --profile=$p attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess --group-name gk8s

# create user
aws iam --profile=$p create-user --user-name cuk8s

# attach user to group
aws iam --profile=$p add-user-to-group --user-name cuk8s --group-name gk8s

aws iam --profile=$p create-access-key --user-name cuk8s
````

# tf examples

# common
terraform init
terraform init -reconfigure
terraform plan
terraform apply
terraform apply -auto-approve
terraform destroy



# sh
cd ed/sh/terraform/examples/sh
# use common commands here

# aws.st
cd ed/sh/terraform/examples/aws.st
# aws lambda
export GOPATH=$PWD/ed/l/go/examples/aws
go get -u "github.com/aws/aws-sdk-go/aws"
# build
GOOS=linux go build -o /tmp/awsLambdaOne $GOPATH/src/app/lambda
cd /tmp && zip awsLambdaOne.zip awsLambdaOne && mv awsLambdaOne.zip /Users/kovpakvolodymyr/Downloads && cd -
# test
open https://realtimelog.herokuapp.com:443/64kfym341kp2
go run $GOPATH/src/app/main.go k31 val200 200
for i in $(seq 2000 2999); do go run $GOPATH/src/app/main.go "k$i" val200 $i; done

# aws.ec2only
cd ed/sh/terraform/examples/aws.ec2only
c=plan
c=apply
c=refresh
c=destroy
terraform $c -var-file=ec2.tfvars -lock=false

# aws.ec2
# @see: docket image build in k8s
cd ed/sh/terraform/examples/aws.ec2/environments/dev
export AWS_PROFILE=x
# use common commands here
# run
h=''
key=/Users/kovpakvolodymyr/web/kovpak/gh/ed/sh/ssh/examples/nopwd/id_rsa
scp -i $key /tmp/xgoapp ec2-user@$h:/tmp
# ssh and run bin file
ssh -i $key ec2-user@$h

# aws.k8s
cd ed/sh/terraform/examples/aws.k8s
# use common commands here
cfg=kubeconfig_my-cluster
d=~/web/kovpak/gh/
kubectl --kubeconfig=$cfg get pods
kubectl --kubeconfig=$cfg apply --force=true -f $d/ed/sh/kubernetes/examples/log/pod.yaml
kubectl --kubeconfig=$cfg apply --force=true -f $d/ed/sh/kubernetes/examples/sh/pod.yaml
kubectl --kubeconfig=$cfg delete pod log-pod
kubectl --kubeconfig=$cfg delete pod ksh-pod



# gcp.storage
cd ed/sh/terraform/examples/gcp.storage
# use common commands here

# gcp.function
cd ed/sh/terraform/examples/gcp.function
zip func.zip func.go
# use common commands here

# gcp.apigee
cd ed/sh/terraform/examples/gcp.apigee/environments/test
